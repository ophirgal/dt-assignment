package forecast

import (
	"log"
	"time"

	"github.com/ophirgal/dt-assignment/backend/internal/config"
	"gorm.io/gorm"
)

func StartWorker(db *gorm.DB, cfg config.Config) {
	interval := time.Duration(cfg.GenerationInterval) * 24 * time.Hour
	go func() {
		for {
			log.Printf("forecast generation attempt: %v", time.Now())
			if err := GenerateForecasts(db, cfg); err != nil {
				log.Printf("forecast generation failed: %v", err)
			}
			time.Sleep(interval)
		}
	}()
}
