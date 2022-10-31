package dao

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"production_service/internal/controller/dto"
)

type ProductStorage struct {
	ID          string
	Name        string
	Description string
	ImageID     sql.NullString
	// TODO следующее видео посвящено работе с деньгами в приложении
	Price         uint64
	CurrencyID    uint32
	Rating        uint32
	CategoryID    uint32
	Specification map[string]interface{}
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

type CreateProductStorageDTO struct {
	ID            string
	Name          string
	Description   string
	ImageID       sql.NullString
	Price         uint64
	CurrencyID    uint32
	Rating        uint32
	CategoryID    uint32
	Specification map[string]interface{}
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

func NewCreateProductStorageDTO(d *dto.CreateProductDTO) *CreateProductStorageDTO {
	now := time.Now()

	var imageID sql.NullString
	if d.ImageID != nil {
		imageID = sql.NullString{
			String: *d.ImageID,
			Valid:  true,
		}
	}
	return &CreateProductStorageDTO{
		ID:            uuid.New().String(),
		Name:          d.Name,
		ImageID:       imageID,
		Description:   d.Description,
		Price:         d.Price,
		CurrencyID:    d.CurrencyID,
		Rating:        d.Rating,
		CategoryID:    d.CategoryID,
		Specification: d.Specification,
		CreatedAt:     now,
		UpdatedAt:     &now,
	}
}
