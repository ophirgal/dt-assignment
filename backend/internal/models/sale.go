package models

import "time"

type Sale struct {
	ID        int       `json:"id"`
	StoreID   int       `json:"storeId"`
	ProductID int       `json:"productId"`
	SoldAt    time.Time `json:"soldAt"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
}
