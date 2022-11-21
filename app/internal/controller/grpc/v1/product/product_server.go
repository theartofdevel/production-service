package product

import (
	"context"

	pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"
	"production_service/internal/domain/product/model"
	"production_service/pkg/logging"
)

func (s *Server) AllProducts(
	ctx context.Context,
	request *pb_prod_products.AllProductsRequest,
) (*pb_prod_products.AllProductsResponse, error) {
	sort := model.ProductsSort(request)
	filter := model.ProductsFilter(request)

	all, err := s.policy.All(ctx, filter, sort)
	if err != nil {
		return nil, err
	}

	pbProducts := make([]*pb_prod_products.Product, len(all))
	for i, p := range all {
		pbProducts[i] = p.ToProto()
	}

	return &pb_prod_products.AllProductsResponse{
		Products: pbProducts,
	}, nil
}

func (s *Server) ProductByID(
	ctx context.Context,
	req *pb_prod_products.ProductByIDRequest,
) (*pb_prod_products.ProductByIDResponse, error) {
	one, err := s.policy.One(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb_prod_products.ProductByIDResponse{
		Product: one.ToProto(),
	}, nil
}

func (s *Server) UpdateProduct(
	ctx context.Context,
	req *pb_prod_products.UpdateProductRequest,
) (*pb_prod_products.UpdateProductResponse, error) {
	product, err := s.policy.One(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	product.UpdateFromPB(req)

	err = s.policy.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return &pb_prod_products.UpdateProductResponse{}, nil
}

func (s *Server) DeleteProduct(
	ctx context.Context,
	req *pb_prod_products.DeleteProductRequest,
) (*pb_prod_products.DeleteProductResponse, error) {
	err := s.policy.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	return &pb_prod_products.DeleteProductResponse{}, nil
}

func (s *Server) CreateProduct(
	ctx context.Context,
	req *pb_prod_products.CreateProductRequest,
) (*pb_prod_products.CreateProductResponse, error) {
	p, err := model.NewProductFromPB(req)
	if err != nil {
		logging.WithError(ctx, err).WithField("product in pb", req).Error("model.NewProductFromPB")
		return nil, err
	}

	product, err := s.policy.CreateProduct(ctx, p)
	if err != nil {
		return nil, err
	}

	return &pb_prod_products.CreateProductResponse{
		Product: product.ToProto(),
	}, nil
}
