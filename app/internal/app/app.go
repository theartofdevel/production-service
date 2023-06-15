package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	_ "production_service/docs"
	"production_service/internal/config"
	"production_service/internal/controller/grpc/v1/product"
	"production_service/internal/domain/product/dao"
	"production_service/internal/domain/product/policy"
	"production_service/internal/domain/product/service"
	"production_service/pkg/common/core/closer"
	"production_service/pkg/common/errors"
	"production_service/pkg/common/logging"
	psql "production_service/pkg/postgresql"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	cfg *config.Config

	router     *httprouter.Router
	httpServer *http.Server
	grpcServer *grpc.Server

	productServiceServer pb_prod_products.ProductServiceServer
}

func NewApp(ctx context.Context, cfg *config.Config) (App, error) {
	logging.L(ctx).Info("router initializing")
	router := httprouter.New()

	logging.L(ctx).Info("swagger docs initializing")
	router.Handler(http.MethodGet, "/swagger", http.RedirectHandler("/swagger/index.html", http.StatusMovedPermanently))
	router.Handler(http.MethodGet, "/swagger/*any", httpSwagger.WrapHandler)

	pgDsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.PostgreSQL.Username,
		cfg.PostgreSQL.Password,
		cfg.PostgreSQL.Host,
		cfg.PostgreSQL.Port,
		cfg.PostgreSQL.Database,
	)

	pgClient, err := psql.NewClient(ctx, 5, 3*time.Second, pgDsn, false)
	if err != nil {
		return App{}, errors.Wrap(err, "psql.NewClient")
	}

	closer.AddN(pgClient)

	productStorage := dao.NewProductStorage(pgClient)
	productService := service.NewProductService(productStorage)
	productPolicy := policy.NewProductPolicy(productService)
	productServiceServer := product.NewServer(
		productPolicy,
		pb_prod_products.UnimplementedProductServiceServer{},
	)

	return App{
		cfg:                  cfg,
		router:               router,
		productServiceServer: productServiceServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startHTTP(ctx)
	})
	grp.Go(func() error {
		return a.startGRPC(ctx, a.productServiceServer)
	})
	return grp.Wait()
}

func (a *App) startGRPC(ctx context.Context, server pb_prod_products.ProductServiceServer) error {
	logger := logging.WithFields(ctx,
		logging.StringField("IP", a.cfg.GRPC.IP),
		logging.IntField("Port", a.cfg.GRPC.Port),
	)

	logger.Info("gRPC Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		logger.With(logging.ErrorField(err)).Fatal("failed to create listener")
	}

	var serverOptions []grpc.ServerOption

	a.grpcServer = grpc.NewServer(serverOptions...)

	pb_prod_products.RegisterProductServiceServer(a.grpcServer, server)

	reflection.Register(a.grpcServer)

	return a.grpcServer.Serve(listener)
}

func (a *App) startHTTP(ctx context.Context) error {
	logger := logging.WithFields(ctx,
		logging.StringField("IP", a.cfg.HTTP.IP),
		logging.IntField("Port", a.cfg.HTTP.Port),
	)
	logger.Info("HTTP Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.HTTP.IP, a.cfg.HTTP.Port))
	if err != nil {
		logger.With(logging.ErrorField(err)).Fatal("failed to create listener")
	}

	logger.With(
		logging.StringsField("AllowedMethods", a.cfg.HTTP.CORS.AllowedMethods),
		logging.StringsField("AllowedOrigins", a.cfg.HTTP.CORS.AllowedOrigins),
		logging.BoolField("AllowCredentials", a.cfg.HTTP.CORS.AllowCredentials),
		logging.StringsField("AllowedHeaders", a.cfg.HTTP.CORS.AllowedHeaders),
		logging.BoolField("OptionsPassthrough", a.cfg.HTTP.CORS.OptionsPassthrough),
		logging.StringsField("ExposedHeaders", a.cfg.HTTP.CORS.ExposedHeaders),
		logging.BoolField("Debug", a.cfg.HTTP.CORS.Debug),
	).Info("CORS initializing")

	c := cors.New(cors.Options{
		AllowedMethods:     a.cfg.HTTP.CORS.AllowedMethods,
		AllowedOrigins:     a.cfg.HTTP.CORS.AllowedOrigins,
		AllowCredentials:   a.cfg.HTTP.CORS.AllowCredentials,
		AllowedHeaders:     a.cfg.HTTP.CORS.AllowedHeaders,
		OptionsPassthrough: a.cfg.HTTP.CORS.OptionsPassthrough,
		ExposedHeaders:     a.cfg.HTTP.CORS.ExposedHeaders,
		Debug:              a.cfg.HTTP.CORS.Debug,
	})

	handler := c.Handler(a.router)

	a.httpServer = &http.Server{
		Handler:      handler,
		WriteTimeout: a.cfg.HTTP.WriteTimeout,
		ReadTimeout:  a.cfg.HTTP.ReadTimeout,
	}

	if err = a.httpServer.Serve(listener); err != nil {
		switch {
		case errors.Is(err, http.ErrServerClosed):
			logger.Warn("server shutdown")
		default:
			logger.With(logging.ErrorField(err)).Fatal("failed to start server")
		}
	}

	err = a.httpServer.Shutdown(context.Background())
	if err != nil {
		logger.With(logging.ErrorField(err)).Fatal("failed to shutdown server")
	}

	return err
}
