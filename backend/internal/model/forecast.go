package model

import (
	"time"

	"gorm.io/gorm"
)

type Forecast struct {
	gorm.Model
	StoreID           uint      `gorm:"not null;uniqueIndex:uniq_forecast;index:idx_forecasts_store_date,priority:1;constraint:OnDelete:CASCADE" json:"storeId"`
	ProductID         uint      `gorm:"not null;uniqueIndex:uniq_forecast;constraint:OnDelete:CASCADE"                                            json:"productId"`
	ForecastDate      time.Time `gorm:"not null;type:date;uniqueIndex:uniq_forecast;index:idx_forecasts_store_date,priority:2"                    json:"forecastDate"`
	Hour              int       `gorm:"not null;check:hour >= 0 AND hour <= 23;uniqueIndex:uniq_forecast"                                        json:"hour"`
	PredictedQuantity float64   `gorm:"not null;type:numeric(10,2)"                                                                              json:"predictedQuantity"`
}

type ForecastResponse struct {
	StoreDisplayName  string  `json:"storeDisplayName"`
	ProductName       string  `json:"productName"`
	ForecastDate      string  `json:"forecastDate"`
	Hour              int     `json:"hour"`
	PredictedQuantity float64 `json:"predictedQuantity"`
}
