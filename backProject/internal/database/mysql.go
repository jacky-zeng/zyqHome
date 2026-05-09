package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"zyqHome/backProject/internal/config"
	"zyqHome/backProject/internal/model"
)

var DB *gorm.DB

func InitMySQL(cfg *config.Config) {
	var err error
	DB, err = gorm.Open(mysql.Open(cfg.DSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect MySQL: %v", err)
	}

	// Auto migrate tables
	err = DB.AutoMigrate(
		&model.User{},
		&model.MenuItem{},
		&model.CenterIcon{},
		&model.SiteConfig{},
		&model.UserBehavior{},
		&model.Image{},
	)
	if err != nil {
		log.Fatalf("Failed to auto migrate: %v", err)
	}

	log.Println("MySQL connected and migrated successfully")
}
