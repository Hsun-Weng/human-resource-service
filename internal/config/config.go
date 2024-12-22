package config

import "os"

const (
	Port             = "PORT"
	DatabaseHost     = "DB_HOST"
	DatabasePort     = "DB_PORT"
	DatabaseUser     = "DB_USER"
	DatabasePassword = "DB_PASS"
	DatabaseName     = "DB_NAME"
	RedisHost        = "REDIS_HOST"
	RedisPort        = "REDIS_PORT"
)

func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
