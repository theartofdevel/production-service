package policy

import (
	"context"

	"production_service/internal/controller/dto"
	"production_service/internal/domain/product/model"
	"production_service/pkg/api/filter"
	"production_service/pkg/api/sort"
	"production_service/pkg/errors"
)

type productService interface {
	All(context.Context, filter.Filterable, sort.Sortable) ([]*model.Product, error)
	Create(context.Context, *dto.CreateProductDTO) (*model.Product, error)
	One(context.Context, string) (*model.Product, error)
}

type ProductPolicy struct {
	productService productService
}

func NewProductPolicy(productService productService) *ProductPolicy {
	return &ProductPolicy{productService: productService}
}

func (p *ProductPolicy) All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]*model.Product, error) {
	products, err := p.productService.All(ctx, filtering, sorting)
	if err != nil {
		return nil, errors.Wrap(err, "productService.All")
	}

	return products, nil
}

func (p *ProductPolicy) CreateProduct(ctx context.Context, d *dto.CreateProductDTO) (*model.Product, error) {
	return p.productService.Create(ctx, d)
}

func (p *ProductPolicy) One(ctx context.Context, id string) (*model.Product, error) {
	return p.productService.One(ctx, id)
}
