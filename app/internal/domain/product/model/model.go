package model

import (
	"encoding/json"
	"time"

	pb_prod_products "github.com/theartofdevel/production-service-contracts/gen/go/prod_service/products/v1"
	"production_service/internal/domain/product/dao"
	"production_service/pkg/logging"
)

type Product struct {
	ID            string
	Name          string
	Description   string
	ImageID       *string
	Price         uint64
	CurrencyID    uint32
	Rating        uint32
	CategoryID    uint32
	Specification map[string]interface{}
	CreatedAt     time.Time
	UpdatedAt     *time.Time
}

func NewProduct(sp *dao.ProductStorage) *Product {
	var imageID *string
	if sp.ImageID.Valid {
		imageID = &sp.ImageID.String
	}

	return &Product{
		ID:            sp.ID,
		Name:          sp.Name,
		Description:   sp.Description,
		ImageID:       imageID,
		Price:         sp.Price,
		CurrencyID:    sp.CurrencyID,
		Rating:        sp.Rating,
		CategoryID:    sp.CategoryID,
		Specification: sp.Specification,
		CreatedAt:     sp.CreatedAt,
		UpdatedAt:     sp.UpdatedAt,
	}
}

func (p *Product) ToProto() *pb_prod_products.Product {
	var updatedAt int64
	if p.UpdatedAt != nil {
		updatedAt = p.UpdatedAt.UnixMilli()
	}

	specBytes, err := json.Marshal(p.Specification)
	if err != nil {
		logging.GetLogger().Warnf("failed to marshal product specification %v", err)
		logging.GetLogger().Trace(p.Specification)
	}

	return &pb_prod_products.Product{
		Id:            p.ID,
		Name:          p.Name,
		Description:   p.Description,
		ImageId:       p.ImageID,
		Price:         p.Price,
		CurrencyId:    p.CurrencyID,
		Rating:        p.Rating,
		CategoryId:    p.CategoryID,
		Specification: string(specBytes),
		UpdatedAt:     updatedAt,
		CreatedAt:     p.CreatedAt.UnixMilli(),
	}
}
