package api

import "github.com/gin-gonic/gin"

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		analytics := v1.Group("/analytics")
		{
			analytics.GET("/stores", GetStores)
			analytics.GET("/sales", GetSales)
		}
	}

	return r
}
