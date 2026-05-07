package dal

import (
	"github.com/ophirgal/dt-assignment/backend/internal/model"
	"gorm.io/gorm"
)

type ForecastRepo struct {
	db *gorm.DB
}

func NewForecastRepo(db *gorm.DB) *ForecastRepo {
	return &ForecastRepo{db: db}
}

func (r *ForecastRepo) GetForecasts(storeID uint, date string) ([]model.ForecastResponse, error) {
	var results []model.ForecastResponse
	err := r.db.
		Model(&model.Forecast{}).
		Select("stores.display_name AS store_display_name, products.name AS product_name, forecasts.forecast_date, forecasts.hour, forecasts.predicted_quantity").
		Joins("JOIN stores ON stores.id = forecasts.store_id").
		Joins("JOIN products ON products.id = forecasts.product_id").
		Where("forecasts.store_id = ? AND forecasts.forecast_date = ?", storeID, date).
		Order("forecasts.hour ASC").
		Scan(&results).Error
	return results, err
}
