package api

import (
	"net/http"

	"dt-assignment/backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetForecasts(c *gin.Context) {
	// TODO: support query params, e.g. forecasts?store=southgate-crossing&date=2026-05-07

	// TODO: add service layer to handle business logic (or use just a simple DAL layer if service layer is not needed?)
	predictions := []models.ForecastResponse{
		// Morning
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Burger",
			ForecastDate:      "2026-05-07",
			Hour:              8,
			PredictedQuantity: 40,
		},
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Fries",
			ForecastDate:      "2026-05-07",
			Hour:              8,
			PredictedQuantity: 60,
		},
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Coke",
			ForecastDate:      "2026-05-07",
			Hour:              8,
			PredictedQuantity: 80,
		},

		// Midday peak
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Burger",
			ForecastDate:      "2026-05-07",
			Hour:              12,
			PredictedQuantity: 120,
		},
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Fries",
			ForecastDate:      "2026-05-07",
			Hour:              12,
			PredictedQuantity: 180,
		},
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Coke",
			ForecastDate:      "2026-05-07",
			Hour:              12,
			PredictedQuantity: 220,
		},

		// Afternoon
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Burger",
			ForecastDate:      "2026-05-07",
			Hour:              16,
			PredictedQuantity: 90,
		},
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Fries",
			ForecastDate:      "2026-05-07",
			Hour:              16,
			PredictedQuantity: 140,
		},
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Coke",
			ForecastDate:      "2026-05-07",
			Hour:              16,
			PredictedQuantity: 170,
		},

		// Evening peak
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Burger",
			ForecastDate:      "2026-05-07",
			Hour:              20,
			PredictedQuantity: 150,
		},
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Fries",
			ForecastDate:      "2026-05-07",
			Hour:              20,
			PredictedQuantity: 220,
		},
		{
			StoreName:         "Southgate Crossing",
			ProductName:       "Coke",
			ForecastDate:      "2026-05-07",
			Hour:              20,
			PredictedQuantity: 260,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": predictions})
}
