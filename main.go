package main

import (
	"nat_backend_go/router"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	dsn := "host=postgres user=user password=secret dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	router.Init(e, db)

	e.Logger.Fatal(e.Start(":3000"))
}
