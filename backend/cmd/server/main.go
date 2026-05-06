package main

import (
	"log"

	"dt-assignment/backend/internal/api"
)

func main() {
	r := api.NewRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
