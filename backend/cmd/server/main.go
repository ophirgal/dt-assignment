package main

import (
	"log"

	"github.com/ophirgal/dt-assignment/backend/internal/api"
	"github.com/ophirgal/dt-assignment/backend/internal/config"
	"github.com/ophirgal/dt-assignment/backend/internal/dal"
	"github.com/ophirgal/dt-assignment/backend/internal/dal/migration"
	"github.com/ophirgal/dt-assignment/backend/internal/forecast"
)

func main() {
	cfg := config.GetConfig()

	db, err := dal.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	log.Println("Running DB migration.")
	if err := migration.Run(db, cfg); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("Starting forecast worker.")
	// Start forecast worker as a separate goroutine.
	// Note: In a real system I consider using a CronJob (K8s), or,
	// if the forecast logic was simple, I would consider using timescaledb's "continuous aggregates".
	forecast.StartWorker(db, cfg)

	log.Println("Starting API server.")
	r := api.NewRouter(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
