package migration

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"sort"
	"time"

	"github.com/ophirgal/dt-assignment/backend/internal/config"
	"github.com/ophirgal/dt-assignment/backend/internal/forecast"
	"github.com/ophirgal/dt-assignment/backend/internal/model"

	"gorm.io/gorm"
)

//go:embed seeds/*.sql
var seedFiles embed.FS

func Run(db *gorm.DB, cfg config.Config) error {
	if err := db.AutoMigrate(
		&model.Chain{},
		&model.Store{},
		&model.Product{},
		&model.Sale{},
		&model.Forecast{},
	); err != nil {
		return fmt.Errorf("automigrate: %w", err)
	}

	entries, err := fs.ReadDir(seedFiles, "seeds")
	if err != nil {
		return fmt.Errorf("read seeds dir: %w", err)
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })

	for _, entry := range entries {
		data, err := seedFiles.ReadFile("seeds/" + entry.Name())
		if err != nil {
			return fmt.Errorf("read seed %s: %w", entry.Name(), err)
		}
		if err := db.Exec(string(data)).Error; err != nil {
			return fmt.Errorf("exec seed %s: %w", entry.Name(), err)
		}
	}

	// backfill forecasts for Jan 2–10, 2026 (seed sales cover Jan 1–9)
	start := time.Date(2026, 1, 2, 0, 0, 0, 0, time.UTC)
	end := time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC)
	for d := start; !d.After(end); d = d.AddDate(0, 0, 1) {
		if err := forecast.GenerateForecastsForDate(db, cfg, d); err != nil {
			log.Printf("backfill forecast %s: %v", d.Format("2006-01-02"), err)
		}
	}

	return nil
}
