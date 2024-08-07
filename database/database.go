package database

import (
	"fmt"
	"log"
	"mini-cms-api/models"
	"mini-cms-api/utils"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var (
	DB_USERNAME string = utils.GetConfig("DB_USERNAME")
	DB_PASSWORD string = utils.GetConfig("DB_PASSWORD")
	DB_NAME     string = utils.GetConfig("DB_NAME")
	DB_HOST     string = utils.GetConfig("DB_HOST")
	DB_PORT     string = utils.GetConfig("DB_PORT")
)

func InitDB() {
	var err error

	var dsn string = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DB_USERNAME,
		DB_PASSWORD,
		DB_HOST,
		DB_PORT,
		DB_NAME,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s\n", err)
	}

	log.Println("connected to the database")

	Migrate()
}

func Migrate() {
	err := DB.AutoMigrate(&models.Category{}, &models.Content{})

	if err != nil {
		log.Fatalf("error when migration: %s\n", err)
	}

	log.Println("migration successful")
}
