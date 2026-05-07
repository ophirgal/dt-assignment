package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ophirgal/dt-assignment/backend/internal/dal"
	"gorm.io/gorm"
)

type ForecastController struct {
	repo *dal.ForecastRepo
}

func NewForecastController(db *gorm.DB) *ForecastController {
	return &ForecastController{repo: dal.NewForecastRepo(db)}
}

func (h *ForecastController) GetForecasts(c *gin.Context) {
	storeID := c.Query("storeId")
	date := c.Query("date")

	if storeID == "" || date == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "storeId and date are required"})
		return
	}

	var id uint
	if _, err := fmt.Sscanf(storeID, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid storeId"})
		return
	}

	forecasts, err := h.repo.GetForecasts(id, date)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": forecasts})
}