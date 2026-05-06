package api

import (
	"net/http"
	"time"

	"dt-assignment/backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetSales(c *gin.Context) {
	now := time.Now().UTC()
	sales := []models.Sale{
		{ID: 1, StoreID: 1, ProductID: 1, SoldAt: now.Add(-2 * time.Hour), Quantity: 3, Total: 3 * 5.99},
		{ID: 2, StoreID: 1, ProductID: 2, SoldAt: now.Add(-1 * time.Hour), Quantity: 5, Total: 5 * 3.49},
		{ID: 3, StoreID: 2, ProductID: 1, SoldAt: now.Add(-30 * time.Minute), Quantity: 2, Total: 2 * 5.99},
	}
	c.JSON(http.StatusOK, gin.H{"data": sales})
}
