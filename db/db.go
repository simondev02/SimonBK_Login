package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "host=localhost user=postgres password=postgres dbname=simon sslmode=disable"
var DBConn *gorm.DB

func ConnectDB() error {
	var err error
	DBConn, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("DB CONNECTED")
	}
	return err
}

func CloseDB() error {
	sqlDB, err := DBConn.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	fmt.Println("DB DISCONNECTED")

	return nil
}
