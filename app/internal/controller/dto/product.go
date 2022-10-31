package dto

import (
	"encoding/json"

	pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"
	"production_service/pkg/logging"
)

type CreateProductDTO struct {
	Name          string
	Description   string
	ImageID       *string
	Price         uint64
	CurrencyID    uint32
	Rating        uint32
	CategoryID    uint32
	Specification map[string]interface{}
}

func NewCreateProductDTOFromPB(product *pb_prod_products.CreateProductRequest) *CreateProductDTO {
	var spec map[string]interface{}
	err := json.Unmarshal([]byte(product.Specification), &spec)
	if err != nil {
		logging.GetLogger().Warnf("failed to unmarshal product specification %v", err)
		logging.GetLogger().Trace(product.Specification)
	}

	return &CreateProductDTO{
		Name:          product.GetName(),
		Description:   product.GetDescription(),
		ImageID:       product.ImageId,
		Price:         product.GetPrice(),
		CurrencyID:    product.GetCurrencyId(),
		Rating:        product.GetRating(),
		CategoryID:    product.GetCategoryId(),
		Specification: spec,
	}
}
