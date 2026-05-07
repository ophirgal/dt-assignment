package model

import (
	"time"

	"gorm.io/gorm"
)

type Store struct {
	gorm.Model
	DisplayName string `gorm:"not null"                                        json:"displayName"`
	SystemName  string `gorm:"not null;uniqueIndex"                             json:"systemName"`
	ChainID     uint   `gorm:"not null;index;constraint:OnDelete:CASCADE"       json:"chainId"`
}

type StoreResponse struct {
	ID          uint           `json:"id"`
	CreatedAt   time.Time      `json:"createdAt"`
	UpdatedAt   time.Time      `json:"updatedAt"`
	DeletedAt   gorm.DeletedAt `json:"deletedAt"`
	DisplayName string         `json:"displayName"`
	SystemName  string         `json:"systemName"`
	ChainID     uint           `json:"chainId"`
}
