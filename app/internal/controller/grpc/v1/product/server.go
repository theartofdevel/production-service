package product

import (
	pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"
	"production_service/internal/domain/policy/product"
)

type Server struct {
	policy *product.Policy
	pb_prod_products.UnimplementedProductServiceServer
}

func NewServer(
	policy *product.Policy,
	srv pb_prod_products.UnimplementedProductServiceServer,
) *Server {
	return &Server{
		policy:                            policy,
		UnimplementedProductServiceServer: srv,
	}
}
