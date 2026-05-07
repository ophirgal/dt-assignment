package config

import (
	"strconv"

	"github.com/joho/godotenv"
	"github.com/ophirgal/dt-assignment/backend/internal/util"
)

type Config struct {
	LookbackDays       int
	GenerationInterval int
}

func GetConfig() Config {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env file")
	}

	lookbackDays := GetPositiveIntFromEnv("LOOKBACK_DAYS", 7)
	generationInterval := GetPositiveIntFromEnv("GENERATION_INTERVAL_DAYS", 1)

	return Config{
		LookbackDays: lookbackDays,
		GenerationInterval: generationInterval,
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
