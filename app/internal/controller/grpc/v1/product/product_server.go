package product

import (
	"context"

	pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"
	"production_service/pkg/common/logging"
)

func (s *Server) AllProducts(
	ctx context.Context,
	request *pb_prod_products.AllProductsRequest,
) (*pb_prod_products.AllProductsResponse, error) {

	return nil, nil
}

func (s *Server) CreateProduct(
	ctx context.Context,
	req *pb_prod_products.CreateProductRequest,
) (*pb_prod_products.CreateProductResponse, error) {
	// mapping and validate
	input, err := NewCreateProductInput(req)
	if err != nil {
		// TODO: response with error code and error message with details
		return nil, err
	}

	// business logic
	createProductOutput, err := s.policy.CreateProduct(ctx, input)
	if err != nil {
		logging.WithError(ctx, err).Error("policy.CreateProduct")
		return nil, err
	}

	response := NewCreateProductResponse(createProductOutput)

	return response, nil
}
