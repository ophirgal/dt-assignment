package forecast

import (
	"fmt"
	"log"

	"github.com/robfig/cron/v3"

	"github.com/ophirgal/dt-assignment/backend/internal/config"
	"gorm.io/gorm"
)

// StartWorker starts the forecast worker as a separate goroutine.
// It relies on the robfig/cron package to schedule the forecast generation task.
// Relying on this package allows us to prevent issues associated with other methods as it uses wall clock (system time).
// (time.Sleep would cause time drift; time.Sleep / time.NewTicker -> would reset the clock if restarted after a crash).
func StartWorker(db *gorm.DB, cfg config.Config) {
	c := cron.New()
	schedule := buildSchedule(cfg)

	_, err := c.AddFunc(schedule, func() {
		log.Printf("forecast generation attempt at scheduled time (hour=%d)", cfg.GenerationHour)
		if err := GenerateForecasts(db, cfg); err != nil {
			log.Printf("forecast generation failed: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("failed to schedule forecast worker: %v", err)
	}

	c.Start()
}

func buildSchedule(cfg config.Config) string {
	schedule := fmt.Sprintf("0 %d */%d * *", cfg.GenerationHour, cfg.GenerationInterval)
	return schedule
}
