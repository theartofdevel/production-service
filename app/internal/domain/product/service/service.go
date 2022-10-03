package service

import (
	"context"

	"production_service/internal/controller/grpc/v1/product"
	"production_service/internal/domain/product/model"
	"production_service/internal/domain/product/storage"
)

type repository interface {
	All(ctx context.Context) ([]storage.Product, error)
	Create(ctx context.Context, dto storage.CreateProductDTO) (storage.Product, error)
}

type Service struct {
	repository repository
}

func (s *Service) All(ctx context.Context) ([]model.Product, error) {

}

func (s *Service) Create(ctx context.Context, dto product.CreateProductDTO) (model.Product, error) {

}
