package config

import (
	"github.com/sisu-network/pairswap-be/src/store"
)

const defaultSisuServeURL = "0.0.0.0:1317"

type AppConfig struct {
	SisuServerURL   string
	SisuGasCostPath string
	Port            int
	DB              store.PostgresConfig
}

func NewDefaultAppConfig() AppConfig {
	// TODO: loads from .env
	dbConfig := store.PostgresConfig{
		Host:     "127.0.0.1",
		Port:     5432,
		Schema:   "pairswap",
		User:     "root",
		Password: "password",
		Option:   "charset=utf8&parseTime=True&loc=Local",
	}

	port := 8080
	return AppConfig{
		SisuServerURL:   defaultSisuServeURL,
		SisuGasCostPath: "/getGasFeeInToken",
		DB:              dbConfig,
		Port:            port,
	}
}
