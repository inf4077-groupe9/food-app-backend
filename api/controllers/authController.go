package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/koybigino/food-app/api/models"
	"github.com/koybigino/food-app/api/oauth2"
	"github.com/koybigino/food-app/api/utils"
	"github.com/koybigino/food-app/api/validations"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(c *fiber.Ctx) error {
	body := new(models.UserLogin)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	errors := validations.ValidateStruct(body)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	user := new(models.User)

	filter := bson.D{
		{Key: "email", Value: body.Email},
	}

	if err := collection.FindOne(context.TODO(), filter).Decode(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	if err := utils.Verify([]byte(body.Password), []byte(user.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	if !user.IsActive {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Unauthorize, Please Validate your email !",
		})
	}

	if user.Token == nil {
		token := oauth2.CreateJWTToken(user.Id, user.UserName, user.Email)
		filter = bson.D{{Key: "username", Value: user.UserName}}

		update := bson.D{{Key: "$set", Value: bson.D{{Key: "token", Value: token}}}}
		_, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		user.Token = token
	}

	return c.JSON(fiber.Map{
		"token":      user.Token,
		"token_type": "Bearer",
	})
}

func Register(c *fiber.Ctx) error {
	body := new(models.UserRegister)

	if err := c.BodyParser(body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": err.Error(),
		})
	}

	errors := validations.ValidateStruct(body)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"errors": errors,
		})
	}

	if body.Password != body.PasswordConfirmation {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Your confirmation password is different to your password !",
		})
	}

	Insertuser := bson.D{
		{Key: "username", Value: body.UserName},
		{Key: "email", Value: body.Email},
		{Key: "password", Value: string(utils.Hash(body.Password))},
		{Key: "token", Value: nil},
		{Key: "is-active", Value: false},
	}

	res, insertErr := collection.InsertOne(context.TODO(), Insertuser)

	if insertErr != nil {
		panic(insertErr)
	}

	user := new(models.User)

	filter := bson.D{
		{Key: "email", Value: body.Email},
	}

	if err := collection.FindOne(context.TODO(), filter).Decode(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "bad Credentials !",
		})
	}

	token := oauth2.CreateJWTToken(user.Id, user.UserName, user.Email)

	filter = bson.D{{Key: "username", Value: body.UserName}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "token", Value: token}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	utils.SendEmail(token, body.Email, body.UserName)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"User":    res,
		"message": "thanks for creating an account, you check your email to validate your email verification !",
	})
}

func EmailVerification(c *fiber.Ctx) error {

	token := c.Params("token")

	filter := bson.D{{Key: "token", Value: token}}

	update := bson.D{{Key: "$set", Value: bson.D{{Key: "is-active", Value: true}}}}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendString("User email verification successfully !")
}
