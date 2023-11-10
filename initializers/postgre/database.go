package postgre

import (
	"NetServDB/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func NewDB(cfg *config.Config) (db *gorm.DB, cleanup func() error, err error) {
	dsn := cfg.DBURL

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
		return nil, nil, err
	}

	log.Print("DB connection successful")

	cleanup = func() error {
		//TODO: погуглить как закрывать горм
		//db.Close()

		return nil
	}

	return db, cleanup, nil
}
