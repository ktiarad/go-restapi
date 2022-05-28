package database

import (
	"fmt"
	"go-restapi/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST = "localhost"
	DB_PORT = 5432
	DB_USER = "ktiarad"
	DB_PASS = "123456"
	DB_NAME = "h8_restapi"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Default().Println("connection db success")

	err = migration(db)
	if err != nil {
		panic(err)
	}

	return db
}

func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(models.Order{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(models.Item{}); err != nil {
		return err
	}
	return nil
}
