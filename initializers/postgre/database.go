package postgre

import (
	"NetServDB/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.DBURL

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
		return nil, err
	}

	log.Print("DB connection successful")

	return db, nil
}
