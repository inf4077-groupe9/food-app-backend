package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/koybigino/food-app/routers"
)

func main() {
	app := fiber.New()

	app.Use(logger.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("hello world !")
	})

	routers.HandleUserRouter(app)
	routers.HandleAuthRoute(app)

	fmt.Println("Server start ...")
	log.Fatal(app.Listen(":3000"))
}
