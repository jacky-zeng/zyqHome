package config

import "os"

type Config struct {
	DBDriver   string
	DBHost     string
	DBPort     string
	DBDatabase string
	DBUsername string
	DBPassword string

	RedisHost string
	RedisAuth string
	RedisPort string

	JWTSecret  string
	ServerPort string
}

func Load() *Config {
	return &Config{
		DBDriver:   getEnv("DB_DRIVER", "mysql"),
		DBHost:     getEnv("DB_HOST", "host.docker.internal"),
		DBPort:     getEnv("DB_PORT", "3306"),
		DBDatabase: getEnv("DB_DATABASE", "ZyqHome"),
		DBUsername: getEnv("DB_USERNAME", "root"),
		DBPassword: getEnv("DB_PASSWORD", "123456"),

		RedisHost: getEnv("REDIS_HOST", "host.docker.internal"),
		RedisAuth: getEnv("REDIS_AUTH", "123456"),
		RedisPort: getEnv("REDIS_PORT", "6379"),

		JWTSecret:  getEnv("JWT_SECRET", "zyqhome-secret-key-2024"),
		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func (c *Config) DSN() string {
	return c.DBUsername + ":" + c.DBPassword + "@tcp(" + c.DBHost + ":" + c.DBPort + ")/" + c.DBDatabase + "?charset=utf8mb4&parseTime=True&loc=Local"
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
