//go:build ignore

// Run with: go run ./internal/forecast/backfill_forecasts.go
// Generates forecasts for Jan 2–10 (inclusive) using sale data from Jan 1–9.
// - uses the default config (LOOKBACK_DAYS=7).

package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/joho/godotenv"
	"github.com/ophirgal/dt-assignment/backend/internal/config"
	"github.com/ophirgal/dt-assignment/backend/internal/dal"
	"github.com/ophirgal/dt-assignment/backend/internal/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	cfg := config.Config{LookbackDays: 7}

	db, err := dal.NewDB()
	if err != nil {
		log.Fatalf("connect db: %v", err)
	}

	// forecast dates: Jan 2 – Jan 10 (inclusive)
	start := time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC)

	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if err := generateForDate(db, cfg, d); err != nil {
			log.Printf("forecast for %s failed: %v", d.Format("2006-01-02"), err)
		} else {
			fmt.Printf("forecast for %s OK\n", d.Format("2006-01-02"))
		}
	}
}

// generateForDate is GenerateForecasts with a fixed forecast date instead of time.Now()+1.
func generateForDate(db *gorm.DB, cfg config.Config, forecastDate time.Time) error {
	// history window: [forecastDate - LookbackDays, forecastDate) — exclusive upper bound includes the day before
	historyEnd := forecastDate
	historyStart := historyEnd.AddDate(0, 0, -cfg.LookbackDays)

	rows, err := computeAverages(db, historyStart, historyEnd)
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

type avgRow struct {
	StoreID   uint
	ProductID uint
	Hour      int
	Avg       float64
}
