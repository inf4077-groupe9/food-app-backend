package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/controllers"
	middleware "github.com/koybigino/food-app/api/middlewares"
)

func HandleUserRouter(router *fiber.App) {
	router.Get("/users/:username", middleware.AuthRequired(), controllers.GetUserByID)
}
