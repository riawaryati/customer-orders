package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "P@ssw0rd"
	dbname   = "orders_by"
)

func ConnectDatabase() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	database, err := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&Item{})
	database.AutoMigrate(&Order{})

	DB = database
}
