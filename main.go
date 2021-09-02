package main

import (
	"go-auth-api/database"
	"go-auth-api/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	database.Connect()

	app := fiber.New()
	app.Use(cors.New(cors.Config{
		// we need AllowCredentials set to true in order to authenticate with HttpOnly Cookie
		AllowCredentials: true,
		AllowOrigins:     "http://localhost:3000",
		AllowMethods:     "GET, POST, DELETE, OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept",
	}))

	routes.Setup(app)

	log.Fatal(app.Listen(":8000"))
}
