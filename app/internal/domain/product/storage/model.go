package storage

import (
	"database/sql"
	"time"

	"production_service/internal/domain/product/model"
)

type Product struct {
	ID            string
	Name          string
	Description   string
	ImageID       sql.NullString
	Price         string
	CurrencyID    uint32
	Rating        uint32
	CategoryID    string
	Specification string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

func (p Product) ToModel() model.Product {
	var imageID *string
	if p.ImageID.Valid {
		imageID = &p.ImageID.String
	}
	return model.Product{
		ID:            p.ID,
		Name:          p.Name,
		Description:   p.Description,
		ImageID:       imageID,
		Price:         p.Price,
		CurrencyID:    p.CurrencyID,
		Rating:        p.Rating,
		CategoryID:    p.CategoryID,
		Specification: p.Specification,
		CreatedAt:     p.CreatedAt,
		UpdatedAt:     p.UpdatedAt,
	}
}

type CreateProductDTO struct {
	Name          string
	Description   string
	ImageID       sql.NullString
	Price         string
	CurrencyID    uint32
	Rating        uint32
	CategoryID    string
	Specification string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}
