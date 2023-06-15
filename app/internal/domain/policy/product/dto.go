package product

import (
	"github.com/shopspring/decimal"
	"production_service/internal/domain/product/model"
	"production_service/pkg/common/errors"
)

type CreateProductInput struct {
	Name          string
	Description   string
	ImageId       *string
	Price         decimal.Decimal
	CurrencyId    uint32
	Rating        uint32
	CategoryId    uint32
	Specification model.Specification
}

func NewCreateProductInput(
	name string,
	description string,
	imageId *string,
	price string,
	currencyId uint32,
	rating uint32,
	categoryId uint32,
) (CreateProductInput, error) {
	input := CreateProductInput{
		Name:          name,
		Description:   description,
		ImageId:       imageId,
		CurrencyId:    currencyId,
		Rating:        rating,
		CategoryId:    categoryId,
		Specification: model.NewSpecification(),
	}

	priceDec, err := decimal.NewFromString(price)
	if err != nil {
		return CreateProductInput{}, errors.Wrap(err, "decimal.NewFromString")
	}

	input.Price = priceDec

	return input, nil
}

type CreateProductOutput struct {
	Product model.Product
}
