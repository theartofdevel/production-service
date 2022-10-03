package product

import pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"

type server struct {
	pb_prod_products.UnimplementedProductServiceServer
}

func NewServer(
	srv pb_prod_products.UnimplementedProductServiceServer,
) *server {
	return &server{UnimplementedProductServiceServer: srv}
}
