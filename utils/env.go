package utils

import "os"

func GetEnv(key string) string {
	return os.Getenv(key)
}

func GetEnvDefault(key string, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		return defaultValue
	}
}
