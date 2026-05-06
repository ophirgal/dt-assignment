package models

import "time"

type Forecast struct {
	ID                int       `json:"id"`
	StoreID           int       `json:"storeId"`
	ProductID         int       `json:"productId"`
	ForecastDate      string    `json:"forecastDate"`
	Hour              int       `json:"hour"`
	PredictedQuantity float64   `json:"predictedQuantity"`
	GeneratedAt       time.Time `json:"generatedAt"`
}
