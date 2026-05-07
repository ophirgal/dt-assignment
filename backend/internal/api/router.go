package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	stores := NewStoreController(db)
	forecasts := NewForecastController(db)

	v1 := r.Group("/api/v1")
	{
		analytics := v1.Group("/analytics")
		{
			analytics.GET("/stores", stores.GetStores)
			analytics.GET("/forecasts", forecasts.GetForecasts)
		}
	}

	return r
}
