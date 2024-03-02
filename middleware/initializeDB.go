package middleware

import (
	"catalogService/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)


func DatabaseConnection() *gorm.DB {
	host := "localhost"
	port := "5432"
	dbName := "postgres"
	dbUser := "postgres"
	password := "postgres"
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, dbUser, dbName, password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database...", err)
	}
	fmt.Println("Database connection successful...")

	db.Exec("DROP TABLE IF EXISTS restaurants;")
	db.Exec("DROP TABLE IF EXISTS menu_items;")

	db.AutoMigrate(&model.MenuItems{})
	db.AutoMigrate(&model.Restaurants{})
	return db
}
