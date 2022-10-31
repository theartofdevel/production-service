package product

import (
	"context"

	pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"
	"production_service/internal/controller/dto"
	"production_service/internal/domain/product/model"
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

func (s *Server) CreateProduct(
	ctx context.Context,
	req *pb_prod_products.CreateProductRequest,
) (*pb_prod_products.CreateProductResponse, error) {
	d := dto.NewCreateProductDTOFromPB(req)

	product, err := s.policy.CreateProduct(ctx, d)
	if err != nil {
		return nil, err
	}

	return &pb_prod_products.CreateProductResponse{
		Product: product.ToProto(),
	}, nil
}
