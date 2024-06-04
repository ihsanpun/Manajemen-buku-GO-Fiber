package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

const DefaultPathAssetImage = "./public/covers/"

func HandleSingleFile(ctx *fiber.Ctx) error {
	//handle file
	file, errFile := ctx.FormFile("cover")

	if errFile != nil {
		log.Println("Error File =", errFile)
	}

	var filename *string
	if file != nil {
		filename = &file.Filename
		//custom nama file
		// extensionFile := filepath.Ext(*filename)
		// newFileName := fmt.Sprintf("gambar-satu%s", extensionFile)
		// errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", newFileName))
		errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", *filename))

		if errSaveFile != nil {
			log.Println("fail to store file into public/covers directory")
		}
	} else {
		log.Println("no file to upload.")
	}
	if filename != nil {
		ctx.Locals("filename", *filename)
	} else {
		ctx.Locals("filename", nil)
	}

	return ctx.Next()

}

func HandleMultipleFile(ctx *fiber.Ctx) error {
	var filenames []string
	form, errForm := ctx.MultipartForm()

	if errForm != nil {
		log.Println("Error Read Multipart Form Request, Error = ", errForm)
	}

	files := form.File["photos"]

	for i, file := range files {
		var filename string
		if file != nil {
			filename = fmt.Sprintf("%d-%s", i, file.Filename)

			errSaveFile := ctx.SaveFile(file, fmt.Sprintf("./public/covers/%s", filename))

			if errSaveFile != nil {
				log.Println("fail to store file into public/covers directory.")
			}
		} else {
			log.Println("no file to upload")
		}
		if filename != "" {
			filenames = append(filenames, filename)
		}
	}
	ctx.Locals("filenames", filenames)
	return ctx.Next()
}

func HandleRemoveFile(filename string, pathFile ...string) error {
	if len(pathFile) > 0 {
		err := os.Remove(pathFile[0] + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	} else {
		err := os.Remove(DefaultPathAssetImage + filename)
		if err != nil {
			log.Println("Failed to remove file")
			return err
		}
	}

	return nil
}
