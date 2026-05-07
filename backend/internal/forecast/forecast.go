package forecast

import (
	"math"
	"time"

	"github.com/ophirgal/dt-assignment/backend/internal/config"
	"github.com/ophirgal/dt-assignment/backend/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type avgRow struct {
	StoreID   uint
	ProductID uint
	Hour      int
	Avg       float64
}

// GenerateForecasts computes next-day forecasts from the last cfg.LookbackDays of history and inserts them.
func GenerateForecasts(db *gorm.DB, cfg config.Config) error {
	today := time.Now().UTC().Truncate(24 * time.Hour)
	start := today.AddDate(0, 0, -cfg.LookbackDays)
	forecastDate := today.AddDate(0, 0, 1)

	rows, err := computeAverages(db, start, today)
	if err != nil {
		return err
	}

	forecasts := make([]model.Forecast, 0, len(rows))
	for _, r := range rows {
		forecasts = append(forecasts, model.Forecast{
			StoreID:           r.StoreID,
			ProductID:         r.ProductID,
			ForecastDate:      forecastDate,
			Hour:              r.Hour,
			PredictedQuantity: math.Ceil(r.Avg),
		})
	}

	if len(forecasts) == 0 {
		return nil
	}

	return db.Clauses(clause.OnConflict{DoNothing: true}).Create(&forecasts).Error
}

// computeAverages returns per-(store, product, hour) average quantity for sales in [start, end).
func computeAverages(db *gorm.DB, start, end time.Time) ([]avgRow, error) {
	var rows []avgRow
	err := db.Raw(`
		SELECT store_id, product_id,
		       EXTRACT(hour FROM sold_at)::int AS hour,
		       AVG(quantity) AS avg
		FROM sales
		WHERE sold_at >= ? AND sold_at < ?
		GROUP BY store_id, product_id, EXTRACT(hour FROM sold_at)
	`, start, end).Scan(&rows).Error
	return rows, err
}
