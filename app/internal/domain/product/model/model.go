package model

import (
	"time"

	pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"
)

type Product struct {
	ID            string
	Name          string
	Description   string
	ImageID       *string
	Price         string
	CurrencyID    uint32
	Rating        uint32
	CategoryID    string
	Specification string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

func (p *Product) ToProto() *pb_prod_products.Product {
	var imageID string
	if p.ImageID != nil {
		imageID = *p.ImageID
	}

	var updatedAt int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt.UnixMilli()
	}

	return &pb_prod_products.Product{
		Id:            p.ID,
		Name:          p.Name,
		Description:   p.Description,
		ImageId:       imageID,
		Price:         p.Price,
		CurrencyId:    p.CurrencyID,
		Rating:        p.Rating,
		CategoryId:    p.CategoryID,
		Specification: p.Specification,
		UpdatedAt:     updatedAt,
		CreatedAt:     p.CreatedAt.UnixMilli(),
	}
}
