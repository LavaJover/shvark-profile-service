package postgres

import (
	"log"

	"github.com/LavaJover/shvark-profile-service/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MustInitDB(cfg *config.ProfileConfig) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.ProfileDB.Dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to init database: %v\n", err)
	}

	db.AutoMigrate(&ProfileModel{})

	return db
}