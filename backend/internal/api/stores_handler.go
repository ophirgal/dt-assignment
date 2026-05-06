package api

import (
	"net/http"

	"dt-assignment/backend/internal/models"

	"github.com/gin-gonic/gin"
)

func GetStores(c *gin.Context) {
	stores := []models.Store{
		{ID: 1, DisplayName: "Downtown Plaza", SystemName: "downtown-plaza"},
		{ID: 2, DisplayName: "North Bridge Mall", SystemName: "north-bridge-mall"},
		{ID: 3, DisplayName: "Riverside Drive", SystemName: "riverside-drive"},
		{ID: 4, DisplayName: "Westgate Terminal", SystemName: "westgate-terminal"},
		{ID: 5, DisplayName: "Eastfield Park", SystemName: "eastfield-park"},
		{ID: 6, DisplayName: "Harborview", SystemName: "harborview"},
		{ID: 7, DisplayName: "Maple & 5th", SystemName: "maple-5th"},
		{ID: 8, DisplayName: "Airport Concourse C", SystemName: "airport-concourse-c"},
		{ID: 9, DisplayName: "University District", SystemName: "university-district"},
		{ID: 10, DisplayName: "Southgate Crossing", SystemName: "southgate-crossing"},
	}
	c.JSON(http.StatusOK, gin.H{"data": stores})
}
