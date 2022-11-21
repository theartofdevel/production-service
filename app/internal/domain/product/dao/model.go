package dao

import (
	"database/sql"
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
	Specification map[string]interface{}
	CreatedAt     sql.NullString
	UpdatedAt     sql.NullString
}
