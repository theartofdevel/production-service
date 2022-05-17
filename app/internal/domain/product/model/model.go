package model

import "time"

type Product struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Description   string     `json:"description"`
	ImageID       *string    `json:"image_id"`
	Price         int        `json:"price"`
	CurrencyID    int        `json:"currency_id"`
	Rating        int        `json:"rating"`
	CategoryID    string     `json:"category_id"`
	Specification string     `json:"specification"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     *time.Time `json:"updated_at"`
}
