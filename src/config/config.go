package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/sisu-network/pairswap-be/src/store"
)

type AppConfig struct {
	SisuServerURL   string
	SisuGasCostPath string
	Port            int
	DB              store.PostgresConfig
}

func NewDefaultAppConfig() AppConfig {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	dbPortStr := os.Getenv("DB_PORT")
	dbPort, err := strconv.Atoi(dbPortStr)
	if err != nil {
		panic(err)
	}

	dbConfig := store.PostgresConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     dbPort,
		Schema:   os.Getenv("DB_SCHEMA"),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Option:   "charset=utf8&parseTime=True&loc=Local",
	}

	portStr := os.Getenv("PORT")
	port, err := strconv.Atoi(portStr)
	if err != nil {
		panic(err)
	}

	sisuServerURL := os.Getenv("SISU_SERVER_URL")
	return AppConfig{
		SisuServerURL:   sisuServerURL,
		SisuGasCostPath: "/getGasFeeInToken",
		DB:              dbConfig,
		Port:            port,
	}
}
