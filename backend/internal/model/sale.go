package model

import (
	"time"

	"gorm.io/gorm"
)

type Sale struct {
	gorm.Model
	StoreID   uint      `gorm:"not null;index:idx_sales_store_sold_at,priority:1;constraint:OnDelete:CASCADE" json:"storeId"`
	ProductID uint      `gorm:"not null;constraint:OnDelete:CASCADE"                                          json:"productId"`
	SoldAt    time.Time `gorm:"not null;index:idx_sales_store_sold_at,priority:2"                             json:"soldAt"`
	Quantity  int       `gorm:"not null"                                                                      json:"quantity"`
	Total     float64   `gorm:"not null;type:numeric(10,2)"                                                   json:"total"`
}
