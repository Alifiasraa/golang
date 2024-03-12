package database

import (
	"fmt"
	"golang-digitalent/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	username = "postgres"
	password = "yeay123"
	port     = 5432
	dbName   = "postgres"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", host, username, password, port, dbName)

	db, err := gorm.Open(postgres.Open(config))
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Order{}, models.Item{})
}