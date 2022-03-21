package store

import (
	"fmt"

	"github.com/sisu-network/lib/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBStores struct {
	TokenStore       *TokenStore
	SupportFormStore *SupportFormStore
	db               *gorm.DB
	Config           PostgresConfig
}

func NewDBStores(cfg PostgresConfig) (*DBStores, error) {
	db, err := ConnectORM(cfg)
	if err != nil {
		return nil, err
	}

	return &DBStores{
		TokenStore:       NewTokenStore(db),
		SupportFormStore: NewSupportFormStore(db),
		db:               db,
		Config:           cfg,
	}, nil
}

func ConnectORM(config PostgresConfig) (*gorm.DB, error) {
	source := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.Host, config.User, config.Password, config.Schema, config.Port)
	log.Info("Database Source: ", source)

	db, err := gorm.Open(postgres.Open(source))
	if err != nil {
		log.Error("cannot connect to database: ", config.Schema)
		return nil, err
	}

	_, err = db.DB()
	if err != nil {
		log.Error("cannot obtain sql database object: ", config.Schema)
		return nil, err
	}

	return db, nil
}
