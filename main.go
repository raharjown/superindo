package main

import (
	"fmt"
	"log"
	"superindo/Database"
	"superindo/Models"
	"superindo/Routes"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	Database.DB, err = gorm.Open("postgres", Database.DBURL(Database.BuildConfig()))
	if err != nil {
		fmt.Println("Status :", err)
	}
	Database.DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`)

	defer Database.DB.Close()
	Database.DB.AutoMigrate(&Models.User{})
	r := Routes.SetupRouter()
	r.Run()
}
