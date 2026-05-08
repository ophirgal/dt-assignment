package config

import (
	"strconv"

	"github.com/joho/godotenv"
	"github.com/ophirgal/dt-assignment/backend/internal/util"
)

// Config holds the application configuration.
// It is loaded from environment variables and has fallback values.
// LOOKBACK_DAYS -> the number of days to look back for forecast generation (integers).
// GENERATION_INTERVAL_DAYS -> the number of days between forecast generations (integers).
// GENERATION_HOUR -> the hour of the day to generate forecasts (0-23).
type Config struct {
	LookbackDays       int
	GenerationInterval int
	GenerationHour     int
}

func GetConfig() Config {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env file")
	}

	lookbackDays := GetPositiveIntFromEnv("LOOKBACK_DAYS", 7)
	generationInterval := GetPositiveIntFromEnv("GENERATION_INTERVAL_DAYS", 1)
	generationHour := GetPositiveIntFromEnv("GENERATION_HOUR", 1)

	return Config{
		LookbackDays:       lookbackDays,
		GenerationInterval: generationInterval,
		GenerationHour:     generationHour,
	}
}

func GetPositiveIntFromEnv(key string, fallback int) int {
	v := util.GetEnv(key, strconv.Itoa(fallback))
	var (
		i   int
		err error
	)
	if i, err = strconv.Atoi(v); err != nil || i <= 0 {
		return fallback
	}

	return i
}
