package models

type ForecastResponse struct {
	StoreName         string  `json:"storeDisplayName"`
	ProductName       string  `json:"productName"`
	ForecastDate      string  `json:"forecastDate"`
	Hour              int     `json:"hour"`
	PredictedQuantity float64 `json:"predictedQuantity"`
}
