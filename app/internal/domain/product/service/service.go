package service

import (
	"context"

	"production_service/internal/domain/product/model"
	"production_service/internal/domain/product/storage"
	"production_service/pkg/api/filter"
	"production_service/pkg/api/sort"
	"production_service/pkg/errors"
)

type repository interface {
	All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]storage.Product, error)
}

type Service struct {
	repository repository
}

func NewProductService(repository repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]model.Product, error) {
	dbProducts, err := s.repository.All(ctx, filtering, sorting)
	if err != nil {
		return nil, errors.Wrap(err, "repository.All")
	}

	var products []model.Product
	for _, dbP := range dbProducts {
		products = append(products, dbP.ToModel())
	}

	return products, nil
}

// func (s *Service) Create(ctx context.Context, dto product.CreateProductDTO) (model.Product, error) {
// 	// TODO продолжить отсюда
// 	return model.Product{}, nil
// }
