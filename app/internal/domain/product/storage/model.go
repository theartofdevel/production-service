package storage

import (
	"database/sql"
	"time"
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
