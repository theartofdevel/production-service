package main

import (
	"context"

	"production_service/internal/app"
	"production_service/internal/config"
	"production_service/pkg/common/logging"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	logging.L(ctx).Info("config initializing")
	cfg := config.GetConfig()

	ctx = logging.ContextWithLogger(ctx, logging.NewLogger())

	a, err := app.NewApp(ctx, cfg)
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.NewApp")
	}

	logging.L(ctx).Info("Running Application")
	err = a.Run(ctx)
	if err != nil {
		logging.WithError(ctx, err).Fatal("app.Run")
		return
	}
}
