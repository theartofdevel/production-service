package model

import (
	"time"

	"github.com/shopspring/decimal"
)

type Product struct {
	ID            string
	Name          string
	Description   string
	ImageID       *string
	Price         decimal.Decimal
	CurrencyID    uint32
	Rating        uint32
	CategoryID    uint32
	Specification Specification
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

func NewProduct(
	ID string,
	name string,
	description string,
	imageID *string,
	price decimal.Decimal,
	currencyID uint32,
	rating uint32,
	categoryID uint32,
	specification Specification,
	createdAt time.Time,
	updatedAt *time.Time,
) Product {
	return Product{
		ID:            ID,
		Name:          name,
		Description:   description,
		ImageID:       imageID,
		Price:         price,
		CurrencyID:    currencyID,
		Rating:        rating,
		CategoryID:    categoryID,
		Specification: specification,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
	}
}

type Specification struct{}

func NewSpecification() Specification {
	return Specification{}
}

type CreateProduct struct {
	ID            string
	Name          string
	Description   string
	ImageID       *string
	Price         decimal.Decimal
	CurrencyID    uint32
	Rating        uint32
	CategoryID    uint32
	Specification Specification
	CreatedAt     time.Time
}

func NewCreateProduct(
	id string,
	name string,
	description string,
	imageID *string,
	price decimal.Decimal,
	currencyID uint32,
	rating uint32,
	categoryID uint32,
	specification Specification,
	createdAt time.Time,
) CreateProduct {
	return CreateProduct{
		ID:            id,
		Name:          name,
		Description:   description,
		ImageID:       imageID,
		Price:         price,
		CurrencyID:    currencyID,
		Rating:        rating,
		CategoryID:    categoryID,
		Specification: specification,
		CreatedAt:     createdAt,
	}
}
