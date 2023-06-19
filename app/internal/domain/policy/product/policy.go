package product

import (
	"context"
	"net/http"
	"time"

	"production_service/internal/apperror"
	"production_service/internal/domain/product/model"
	"production_service/internal/domain/product/service"
	"production_service/pkg/common/core/clock"
	"production_service/pkg/common/errors"
)

var (
	ErrProductPriceIsNegative = apperror.NewAppError(http.StatusBadRequest, "00250", "product price is negative")
	ErrProductPriceIsZero     = apperror.NewAppError(http.StatusBadRequest, "00251", "product price is zero")
	ErrProductRating          = apperror.NewAppError(http.StatusBadRequest, "00252", "invalid product rating")
)

type IdentityGenerator interface {
	GenerateUUIDv4String() string
}

type Clock interface {
	Now() time.Time
}

type Policy struct {
	productService *service.ProductService

	identity IdentityGenerator
	clock    Clock
}

func NewProductPolicy(productService *service.ProductService, identity IdentityGenerator, clock clock.Clock) *Policy {
	return &Policy{
		productService: productService,

		identity: identity,
		clock:    clock,
	}
}

func (p *Policy) All(ctx context.Context) ([]model.Product, error) {
	products, err := p.productService.All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "productService.All")
	}

	return products, nil
}

func (p *Policy) CreateProduct(ctx context.Context, input CreateProductInput) (CreateProductOutput, error) {
	if input.Rating < 0 || input.Rating > 5 {
		return CreateProductOutput{}, ErrProductRating
	}
	if input.Price.IsNegative() {
		return CreateProductOutput{}, ErrProductPriceIsNegative
	}

	if input.Price.IsZero() {
		return CreateProductOutput{}, ErrProductPriceIsZero
	}

	createProduct := model.NewCreateProduct(
		p.identity.GenerateUUIDv4String(),
		input.Name,
		input.Description,
		input.ImageId,
		input.Price,
		input.CurrencyId,
		input.Rating,
		input.CategoryId,
		input.Specification,
		p.clock.Now(),
	)

	product, err := p.productService.CreateProduct(ctx, createProduct)
	if err != nil {
		return CreateProductOutput{}, errors.Wrap(err, "productService.CreateProduct")
	}

	return CreateProductOutput{
		Product: product,
	}, nil
}
