package db

import (
	"go-rest-api-halil-cin/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := "host=db" +
		" user=" + cfg.PGUser +
		" password=" + cfg.PGPassword +
		" dbname=" + cfg.PGName +
		" port=" + cfg.PGPort +
		" sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
