package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ophirgal/dt-assignment/backend/internal/dal"
	"gorm.io/gorm"
)

type StoreController struct {
	repo *dal.StoreRepo
}

func NewStoreController(db *gorm.DB) *StoreController {
	return &StoreController{repo: dal.NewStoreRepo(db)}
}

func (h *StoreController) GetStores(c *gin.Context) {
	stores, err := h.repo.GetStores()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": stores})
}
