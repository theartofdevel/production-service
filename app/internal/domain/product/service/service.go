package service

import (
	"context"

	"production_service/internal/domain/product/dao"
	"production_service/internal/domain/product/model"
	"production_service/pkg/api/filter"
	"production_service/pkg/api/sort"
	"production_service/pkg/errors"
)

type repository interface {
	All(context.Context, filter.Filterable, sort.Sortable) ([]*dao.ProductStorage, error)
	One(context.Context, string) (*dao.ProductStorage, error)
	Create(context.Context, map[string]interface{}) error
	Update(context.Context, string, map[string]interface{}) error
	Delete(context.Context, string) error
}

type ProductService struct {
	repository repository
}

func NewProductService(repository repository) *ProductService {
	return &ProductService{repository: repository}
}

func (s *ProductService) All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]*model.Product, error) {
	dbProducts, err := s.repository.All(ctx, filtering, sorting)
	if err != nil {
		return nil, errors.Wrap(err, "repository.All")
	}

	var products []*model.Product
	for _, dbP := range dbProducts {
		products = append(products, model.NewProduct(dbP))
	}

	return products, nil
}

func (s *ProductService) Create(ctx context.Context, product *model.Product) (*model.Product, error) {
	productStorageMap, err := product.ToMap()
	if err != nil {
		return nil, err
	}

	err = s.repository.Create(ctx, productStorageMap)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) One(ctx context.Context, id string) (*model.Product, error) {
	one, err := s.repository.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.NewProduct(one), nil
}

func (s *ProductService) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *ProductService) Update(ctx context.Context, product *model.Product) error {
	productStorageMap, err := product.ToMap()
	if err != nil {
		return err
	}

	return s.repository.Update(ctx, product.ID, productStorageMap)
}
