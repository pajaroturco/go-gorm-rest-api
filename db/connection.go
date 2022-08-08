package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var dsn = "host=localhost user=pajaro password=mysecretpassword dbname=admin port=5432 sslmode=disable TimeZone=America/Mexico_City"
var DB *gorm.DB

func DBConnection() {

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB connected")
	}

}
