package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	stringConnection := "host=localhost user=postgres password=root dbname=api_personalidades port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil{
		log.Panic(err)
	}
}