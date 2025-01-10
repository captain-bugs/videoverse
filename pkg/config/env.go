package config

import (
	"os"
	"strconv"
)

func getEnvString(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		value = defaultValue
	}
	return value
}

func getEnvInt(key string, defaultValue int) int {
	envValue := os.Getenv(key)
	if len(envValue) == 0 {
		value := defaultValue
		return value
	}
	value, err := strconv.Atoi(envValue)
	if err != nil {
		return defaultValue
	}
	return value
}
