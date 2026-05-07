package dal

import (
	"github.com/ophirgal/dt-assignment/backend/internal/model"
	"gorm.io/gorm"
)

type StoreRepo struct {
	db *gorm.DB
}

func NewStoreRepo(db *gorm.DB) *StoreRepo {
	return &StoreRepo{db: db}
}

func (r *StoreRepo) GetStores() ([]model.StoreResponse, error) {
	var stores []model.Store
	if err := r.db.Find(&stores).Error; err != nil {
		return nil, err
	}

	var response []model.StoreResponse
	for _, store := range stores {
		response = append(response, model.StoreResponse{
			ID:          store.ID,
			CreatedAt:   store.CreatedAt,
			UpdatedAt:   store.UpdatedAt,
			DeletedAt:   store.DeletedAt,
			DisplayName: store.DisplayName,
			SystemName:  store.SystemName,
			ChainID:     store.ChainID,
		})
	}

	return response, nil
}
