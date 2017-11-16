package config

import (
	"os"
	"strconv"
)

// AppConfig is a struct for storing the app config
type AppConfig struct {
	Port int
}

var App AppConfig

var version string

func init() {
	App = AppConfig{
		Port: envInt("PORT", 9000),
	}
}

func envString(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func envBool(key string, defaultValue bool) bool {
	valueStr := os.Getenv(key)
	if valueStr == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func envInt(key string, defaultValue int) int {
	strValue := os.Getenv(key)
	if strValue == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(strValue)
	if err != nil {
		panic(err)
	}
	return value
}
