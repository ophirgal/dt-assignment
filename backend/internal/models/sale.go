package models

import "time"

type Sale struct {
	ID        int       `json:"id"`
	StoreID   int       `json:"store_id"`
	ProductID int       `json:"product_id"`
	SoldAt    time.Time `json:"sold_at"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
}
