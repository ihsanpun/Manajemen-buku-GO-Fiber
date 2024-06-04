package handler

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ihsanpun/go-fiber-part2/database"
	"github.com/ihsanpun/go-fiber-part2/model/entity"
	"github.com/ihsanpun/go-fiber-part2/model/request"
)

func BookHandlerRead(ctx *fiber.Ctx) error {
	// userInfo := ctx.Locals("userInfo")
	// log.Println("user info data ::", userInfo)

	var books []entity.Book
	result := database.DB.Debug().Find(&books)
	if result.Error != nil {
		log.Println(result.Error)
	}

	return ctx.JSON(books)
}

func BookHandlerCreate(ctx *fiber.Ctx) error {
	book := new(request.BookCreateRequest)

	if err := ctx.BodyParser(book); err != nil {
		return err
	}

	//validation
	validate := validator.New()
	errValidate := validate.Struct(book)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// validation required image
	filename := ctx.Locals("filename")
	var filenameString string
	log.Println("filename = ", filename)
	if filename == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "cover image required",
		})
	} else {
		filenameString = fmt.Sprintf("%v", filename)
	}

	newBook := entity.Book{
		Title:  book.Title,
		Author: book.Author,
		Cover:  filenameString,
	}

	errCreateBook := database.DB.Create(&newBook).Error
	if errCreateBook != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "failed when creating data",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
		"data":    newBook,
	})
}

func BookHandlerGetById(ctx *fiber.Ctx) error {
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

func BookHandlerUpdate(ctx *fiber.Ctx) error {
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

func BookHandlerDelete(ctx *fiber.Ctx) error {
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
