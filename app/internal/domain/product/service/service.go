package service

import (
	"context"

	"production_service/internal/domain/product/model"
	"production_service/pkg/common/errors"
)

type repository interface {
	All(ctx context.Context) ([]model.Product, error)
	Create(ctx context.Context, req model.CreateProduct) error
}

type ProductService struct {
	repository repository
}

func NewProductService(repository repository) *ProductService {
	return &ProductService{repository: repository}
}

func (s *ProductService) All(ctx context.Context) ([]model.Product, error) {
	products, err := s.repository.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "repository.All")
	}

	return products, nil
}

func (s *ProductService) CreateProduct(ctx context.Context, req model.CreateProduct) (model.Product, error) {
	// cache

	err := s.repository.Create(ctx, req)
	if err != nil {
		return model.Product{}, err
	}

	return model.NewProduct(
		req.ID,
		req.Name,
		req.Description,
		req.ImageID,
		req.Price,
		req.CurrencyID,
		req.Rating,
		req.CategoryID,
		req.Specification,
		req.CreatedAt,
		nil,
	), nil
}
