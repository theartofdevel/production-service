package policy

import (
	"context"

	"production_service/internal/domain/product/model"
	"production_service/pkg/api/filter"
	"production_service/pkg/api/sort"
	"production_service/pkg/errors"
)

type productService interface {
	All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.Product, error)
	// Create(ctx context.Context, dto product.CreateProductDTO) (model.Product, error)
}

type ProductPolicy struct {
	productService productService
}

func NewProductPolicy(productService productService) *ProductPolicy {
	return &ProductPolicy{productService: productService}
}

func (p *ProductPolicy) All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.Product, error) {
	products, err := p.productService.All(ctx, filtering, sorting)
	if err != nil {
		return nil, errors.Wrap(err, "productService.All")
	}

	return products, nil
}
