package controllers

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
	"github.com/koybigino/food-app/database"
	"go.mongodb.org/mongo-driver/bson"
)

var client = database.Connection()
var collection = client.Database(os.Getenv("DB_NAME")).Collection("users")

func GetUserByID(c *fiber.Ctx) error {
	usernameParams := c.Params("username")

	usernameSplit := strings.Split(usernameParams, "-")

	username := strings.Join(usernameSplit, " ")

	user := new(models.User)
	userResponse := new(models.UserResponse)

	filter := bson.D{
		{Key: "username", Value: username},
	}

	if err := collection.FindOne(context.TODO(), filter).Decode(user); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error":   err.Error(),
			"message": fmt.Sprintf("Any User correspond to this username = %s", username),
		})
	}

	models.ParseToUserResponse(*user, userResponse)

	return c.JSON(fiber.Map{
		"User": userResponse,
	})
}
