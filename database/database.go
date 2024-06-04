package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=fibergo port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic("cannot connect to database")
	}

	fmt.Println("connected to database")
}
