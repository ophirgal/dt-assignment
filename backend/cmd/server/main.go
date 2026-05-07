package main

import (
	"log"

	"github.com/ophirgal/dt-assignment/backend/internal/api"
	"github.com/ophirgal/dt-assignment/backend/internal/dal"
	"github.com/ophirgal/dt-assignment/backend/internal/dal/migration"
)

func main() {
	db, err := dal.NewDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	if err := migration.Run(db); err != nil {
		log.Fatalf("migration failed: %v", err)
	}

	r := api.NewRouter(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
