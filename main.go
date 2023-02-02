package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/isaiorellana-dev/fiber/routes"
)

func main() {
	app := fiber.New()

	routes.UseMoviesRoutes(app)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://gofiber.io, http://localhost:3000/",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Listen(":4000")

}
