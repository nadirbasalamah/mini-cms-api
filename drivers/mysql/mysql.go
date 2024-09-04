package mysql_driver

import (
	"fmt"
	"log"
	"mini-cms-api/drivers/mysql/categories"
	"mini-cms-api/drivers/mysql/contents"
	"mini-cms-api/drivers/mysql/users"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *DBConfig) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	pgConfig := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(pgConfig, &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s\n", err)
	}

	log.Println("connected to the database")

	return db
}

func MigrateDB(db *gorm.DB) {
	err := db.AutoMigrate(&categories.Category{}, &contents.Content{}, &users.User{})

	if err != nil {
		log.Fatalf("database migration failed: %v\n", err)
	}
}
