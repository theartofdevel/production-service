package product

import pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"

type CreateProductDTO struct {
	Name          string
	Description   string
	ImageID       *string
	Price         string
	CurrencyID    uint32
	Rating        uint32
	CategoryID    string
	Specification string
}

func NewCreateProductDTOFromPB(product *pb_prod_products.CreateProductRequest) *CreateProductDTO {
	return &CreateProductDTO{
		Name:          product.GetName(),
		Description:   product.GetDescription(),
		ImageID:       product.ImageId,
		Price:         product.GetPrice(),
		CurrencyID:    product.GetCurrencyId(),
		Rating:        product.GetRating(),
		CategoryID:    product.GetCategoryId(),
		Specification: product.GetSpecification(),
	}
}
