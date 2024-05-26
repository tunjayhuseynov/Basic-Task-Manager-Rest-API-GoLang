package main

import (
	"crud-restapi/router"
	"crud-restapi/database"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var App = fiber.New()

func main() {
	envErr := godotenv.Load(".env")
	db, err := database.InitDatabase()

	router.Router(App, db)
	
	
	if envErr != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}


	App.Listen(":3000")
}
