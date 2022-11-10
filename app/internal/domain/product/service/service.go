package service

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"production_service/internal/controller/dto"
	"production_service/internal/domain/product/dao"
	"production_service/internal/domain/product/model"
	"production_service/pkg/api/filter"
	"production_service/pkg/api/sort"
	"production_service/pkg/errors"
)

type repository interface {
	All(context.Context, filter.Filterable, sort.Sortable) ([]*dao.ProductStorage, error)
	One(context.Context, string) (*dao.ProductStorage, error)
	Create(context.Context, *dao.CreateProductStorageDTO) error
	Delete(context.Context, string) error
	Update(context.Context, string, map[string]interface{}) error
}

type Service struct {
	repository repository
}

func NewProductService(repository repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) All(ctx context.Context, filtering filter.Filterable, sorting sort.Sortable) ([]*model.Product, error) {
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

func (s *Service) Create(ctx context.Context, d *dto.CreateProductDTO) (*model.Product, error) {
	createProductStorageDTO := dao.NewCreateProductStorageDTO(d)

	err := s.repository.Create(ctx, createProductStorageDTO)
	if err != nil {
		return nil, err
	}

	one, err := s.repository.One(ctx, createProductStorageDTO.ID)
	if err != nil {
		return nil, err
	}

	return model.NewProduct(one), nil
}

func (s *Service) One(ctx context.Context, id string) (*model.Product, error) {
	one, err := s.repository.One(ctx, id)
	if err != nil {
		return nil, err
	}

	return model.NewProduct(one), nil
}

func (s *Service) Delete(ctx context.Context, id string) error {
	return s.repository.Delete(ctx, id)
}

func (s *Service) Update(ctx context.Context, id string, d *dto.UpdateProductDTO) error {
	storageDTO := dao.NewUpdateProductStorageDTO(d)

	var updateProductMap map[string]interface{}

	err := mapstructure.Decode(storageDTO, &updateProductMap)
	if err != nil {
		return errors.Wrap(err, "mapstructure.Decode UpdateProductDTO")
	}

	return s.repository.Update(ctx, id, updateProductMap)
}
