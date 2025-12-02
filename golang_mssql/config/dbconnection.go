package config

import (
	"fmt"
	"log"
	"src/golang_mssql/models"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() *gorm.DB {
	var err error
	dsn := "sqlserver://sa:Reynald88@127.0.0.1:1433?database=golang_angular20"
	// dsn := "server=localhost;user id=sa;password=Reynald88;database=golang_angular20;TrustServerCertificate=False"
	DB, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Could not connect to MSSQL SERVER 2019 database.")
	} else {
		fmt.Println("connected to mssql server 2019 database.")
	}

	DB.AutoMigrate(&models.User{}, &models.Product{})
	log.Print("Tables Created....")

	return DB
}
