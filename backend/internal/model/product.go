package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name    string  `gorm:"not null"                                        json:"name"`
	Price   float64 `gorm:"not null;type:numeric(10,2)"                     json:"price"`
	ChainID uint    `gorm:"not null;index;constraint:OnDelete:CASCADE"      json:"chainId"`
}
