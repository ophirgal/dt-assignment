package dal

import (
	"fmt"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ophirgal/dt-assignment/backend/internal/util"
)

func NewDB() (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		util.GetEnv("DB_HOST", "localhost"),
		util.GetEnv("DB_PORT", "5432"),
		util.GetEnv("DB_USER", "postgres"),
		util.GetEnv("DB_PASSWORD", "postgres"),
		util.GetEnv("DB_NAME", "postgres"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// connection pool settings
	sqlDB.SetMaxOpenConns(25)
	sqlDB.SetMaxIdleConns(25)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	return db, nil
}
