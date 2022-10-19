package product

import (
	pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"
	"production_service/internal/domain/product/policy"
)

type Server struct {
	policy *policy.ProductPolicy
	pb_prod_products.UnimplementedProductServiceServer
}

func NewServer(
	policy *policy.ProductPolicy,
	srv pb_prod_products.UnimplementedProductServiceServer,
) *Server {
	return &Server{
		policy:                            policy,
		UnimplementedProductServiceServer: srv,
	}
}
