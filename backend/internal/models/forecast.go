package models

import "time"

type Forecast struct {
	ID                int       `json:"id"`
	StoreID           int       `json:"store_id"`
	ProductID         int       `json:"product_id"`
	ForecastDate      string    `json:"forecast_date"`
	Hour              int       `json:"hour"`
	PredictedQuantity float64   `json:"predicted_quantity"`
	GeneratedAt       time.Time `json:"generated_at"`
}
