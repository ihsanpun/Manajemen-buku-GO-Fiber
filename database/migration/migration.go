package migration

import (
	"fmt"
	"log"

	"github.com/ihsanpun/go-fiber-part2/database"
	"github.com/ihsanpun/go-fiber-part2/model/entity"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&entity.User{}, &entity.Book{}, &entity.Category{}, &entity.Photo{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}
