package policy

import (
	"context"

	"production_service/internal/controller/grpc/v1/product"
	"production_service/internal/domain/product/model"
)

type productService interface {
	All(ctx context.Context) ([]model.Product, error)
	Create(ctx context.Context, dto product.CreateProductDTO) (model.Product, error)
}

type ProductPolicy struct {
	productService productService
}
