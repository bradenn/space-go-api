package Config

import (
	"os"
	"strconv"
)

type EnvVars struct {
	Port   int
	DBPort int
	DBHost string
	DBUser string
	DBName string
	DBPass string
}

func GetConfig() EnvVars {
	return EnvVars{
		Port:   getEnvAsInt("PORT", 8080),
		DBPort: getEnvAsInt("DB_PORT", 5432),
		DBHost: getEnv("DB_HOST", "host"),
		DBUser: getEnv("DB_USER", "user"),
		DBPass: getEnv("DB_PASS", "pass"),
		DBName: getEnv("DB_NAME", "name"),
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
