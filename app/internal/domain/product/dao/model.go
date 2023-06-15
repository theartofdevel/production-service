package dao

import (
	"database/sql"
	"time"

	"production_service/internal/domain/product/model"
	"production_service/pkg/utils/pointer"
)

type ProductStorage struct {
	ID            string
	Name          string
	Description   string
	ImageID       sql.NullString
	Price         uint64
	CurrencyID    uint32
	Rating        uint32
	CategoryID    uint32
	Specification Specification
	CreatedAt     time.Time
	UpdatedAt     sql.NullTime
}

type Specification struct{}

func (s *Specification) ToDomain() model.Specification {
	return model.Specification{}
}

func (ps *ProductStorage) ToDomain() model.Product {
	var ImageID *string
	if ps.ImageID.Valid {
		ImageID = pointer.Pointer(ps.ImageID.String)
	}

	var UpdatedAt *time.Time
	if ps.UpdatedAt.Valid {
		UpdatedAt = pointer.Pointer(ps.UpdatedAt.Time)
	}

	return model.Product{
		ID:            ps.ID,
		Name:          ps.Name,
		Description:   ps.Description,
		ImageID:       ImageID,
		Price:         ps.Price,
		CurrencyID:    ps.CurrencyID,
		Rating:        ps.Rating,
		CategoryID:    ps.CategoryID,
		Specification: ps.Specification.ToDomain(),
		CreatedAt:     ps.CreatedAt,
		UpdatedAt:     UpdatedAt,
	}
}
