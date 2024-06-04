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

func PhotoHandlerCreate(ctx *fiber.Ctx) error {
	photo := new(request.PhotoCreateRequest)

	if err := ctx.BodyParser(photo); err != nil {
		return err
	}

	//validation
	validate := validator.New()
	errValidate := validate.Struct(photo)
	if errValidate != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	// validation required image
	filenames := ctx.Locals("filenames")

	if filenames == nil {
		return ctx.Status(422).JSON(fiber.Map{
			"message": "cover image required",
		})
	} else {
		filenamesData := filenames.([]string)

		for _, filename := range filenamesData {
			newPhoto := entity.Photo{
				Image:      filename,
				CategoryId: photo.CategoryId,
			}

			errCreatePhoto := database.DB.Create(&newPhoto).Error
			if errCreatePhoto != nil {
				log.Println("some photo cannot uploaded properly")
			}
		}
	}

	return ctx.JSON(fiber.Map{
		"message": "success",
	})
}

func PhotoHandlerDelete(ctx *fiber.Ctx) error {
	photoId := ctx.Params("id")
	var photo entity.Photo

	errFindPhoto := database.DB.First(&photo, "id = ?", photoId).Error

	if errFindPhoto != nil {
		ctx.Status(404).JSON(fiber.Map{
			"message": "Photo not found",
		})
	}
	//Handler delete covers
	utils.HandleRemoveFile(photo.Image)

	//Delete Database
	errDelete := database.DB.Debug().Delete(&photo).Error

	if errDelete != nil {
		ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "Photo deleted succesfully",
	})
}
