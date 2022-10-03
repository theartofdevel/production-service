package model

import "time"

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
