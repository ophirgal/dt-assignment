package model

import "gorm.io/gorm"

type Chain struct {
	gorm.Model
	Name string `gorm:"not null;uniqueIndex" json:"name"`
}
