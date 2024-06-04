package handler

import (
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ihsanpun/go-fiber-part2/database"
	"github.com/ihsanpun/go-fiber-part2/model/entity"
	"github.com/ihsanpun/go-fiber-part2/model/request"
	"github.com/ihsanpun/go-fiber-part2/utils"
)

func UserHandlerRead(ctx *fiber.Ctx) error {
	userInfo := ctx.Locals("userInfo")
	log.Println("user info data ::", userInfo)

	var users []entity.User
	result := database.DB.Debug().Find(&users)
	if result.Error != nil {
		log.Println(result.Error)
	}

	//kode ini sama aja kaya yang atas
	// err := database.DB.Find(&users).Error
	// if err != nil {
	// 	log.Println(err)
	// }

	return ctx.JSON(users)
}

func UserHandlerCreate(ctx *fiber.Ctx) error {
	user := new(request.UserCreateRequest)

	if err := ctx.BodyParser(user); err != nil {
		return err
	}

	validate := validator.New()
	errValidate := validate.Struct(user)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Phone:   user.Phone,
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newUser.Password = hashedPassword

	errCreateUser := database.DB.Create(&newUser).Error
	if errCreateUser != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed when creating data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newUser,
	})
}

func UserHandlerGetById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user entity.User
	// var user response.UserResponse
	err := database.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	// userResponse := response.UserResponse{
	// 	ID:        user.ID,
	// 	Name:      user.Name,
	// 	Email:     user.Email,
	// 	Address:   user.Address,
	// 	Phone:     user.Phone,
	// 	CreatedAt: user.CreatedAt,
	// 	UpdatedAt: user.UpdatedAt,
	// }

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerUpdate(ctx *fiber.Ctx) error {
	//bind request data ke variabel userRequest
	userRequest := new(request.UserUpdateRequest)
	if err := ctx.BodyParser(userRequest); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "bad request",
		})
	}
	//carii user
	id := ctx.Params("id")
	var user entity.User
	//cek user ada apa ngga
	err := database.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	//update data
	if userRequest.Name != "" {
		user.Name = userRequest.Name
	}
	user.Address = userRequest.Address
	user.Phone = userRequest.Phone
	errUpdate := database.DB.Save(&user).Error

	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    user,
	})
}

func UserHandlerDelete(ctx *fiber.Ctx) error {
	userId := ctx.Params("id")
	var user entity.User

	errFindUser := database.DB.First(&user, "id = ?", userId).Error

	if errFindUser != nil {
		ctx.Status(404).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	errDelete := database.DB.Debug().Delete(&user).Error

	if errDelete != nil {
		ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "user deleted succesfully",
	})

}
