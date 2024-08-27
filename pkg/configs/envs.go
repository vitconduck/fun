package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		APP  *APP
		HTTP *HTTP
		DB   *DB
	}

	APP struct {
		Name string
		Env  string
	}

	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}

	DB struct {
		PublicHost             string
		Port                   string
		DBUser                 string
		DBPassword             string
		DBAddress              string
		DBName                 string
		JWTSecret              string
		JWTExpirationInSeconds int64
	}
)

func New() (*Config, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	db := DB{
		DBUser:                 getEnv("DB_USER", "postgres"),
		DBPassword:             getEnv("DB_PASSWORD", "mypassword"),
		DBAddress:              fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "5432")),
		DBName:                 getEnv("DB_NAME", "fun"),
		JWTSecret:              getEnv("JWT_SECRET", "not-so-secret-now-is-it?"),
		JWTExpirationInSeconds: getEnvAsInt("JWT_EXPIRATION_IN_SECONDS", 3600*24*7),
	}

	app := APP{
		Name: getEnv("APP_NAME", "fun"),
		Env:  getEnv("APP_ENV", "dev"),
	}

	http := HTTP{
		Env:            getEnv("APP_ENV", "development"),
		URL:            getEnv("HTTP_URL", "127.0.0.1"),
		Port:           getEnv("HTTP_PORT", "8080"),
		AllowedOrigins: getEnv("HTTP_ALLOWED_ORIGINS", "http://127.0.0.1:3000,http://127.0.0.1:5173"),
	}

	return &Config{
		APP:  &app,
		HTTP: &http,
		DB:   &db,
	}, nil
}

// Gets the env by key or fallbacks
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return fallback
}

func getEnvAsInt(key string, fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fallback
		}

		return i
	}

	return fallback
}
