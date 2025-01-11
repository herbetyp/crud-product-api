package database

import (
	"fmt"

	"github.com/herbetyp/crud-product-api/configs"
	"github.com/herbetyp/crud-product-api/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDatabase() {
	dbConf := configs.GetConfig().DB

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	dbConf.Host, dbConf.Port, dbConf.User, dbConf.Password, dbConf.DBName, dbConf.SSLmode)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Could not connect to the Postgres Database")
	}

	db = database

	config, _ := database.DB()
	config.SetMaxIdleConns(dbConf.SetMaxIdleConns)
	config.SetMaxOpenConns(dbConf.SetMaxOpenConns)
	config.SetConnMaxLifetime(dbConf.SetConnMaxLifetime)

	
	fmt.Printf("Connected to Postgres Database")
	migrations.AutoMigrations(db)
}

func GetDatabase() *gorm.DB {
	return db
}
