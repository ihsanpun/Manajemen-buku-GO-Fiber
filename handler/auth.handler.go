package handler

import (
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ihsanpun/go-fiber-part2/database"
	"github.com/ihsanpun/go-fiber-part2/model/entity"
	"github.com/ihsanpun/go-fiber-part2/model/request"
	"github.com/ihsanpun/go-fiber-part2/utils"
)

func LoginHandler(ctx *fiber.Ctx) error {
	loginRequest := new(request.LoginRequest)

	if err := ctx.BodyParser(loginRequest); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(loginRequest)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	//check available user
	var user entity.User

	err := database.DB.First(&user, "email = ?", loginRequest.Email).Error
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Credentials",
		})
	}

	//check validation password

	isValid := utils.CheckPasswordHash(loginRequest.Password, user.Password)
	if !isValid {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Credentials",
		})
	}

	//Generate JWT
	claims := jwt.MapClaims{}
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["address"] = user.Address
	claims["exp"] = time.Now().Add(time.Minute * 10).Unix() //Token expired dalam 2 menit

	if user.Email == "john@gmail.com" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "user"
	}

	token, errGenerateToken := utils.GenerateToken(&claims)

	if errGenerateToken != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "Wrong Credentials",
		})
	}

	return ctx.JSON(fiber.Map{
		"token": token,
	})
}
