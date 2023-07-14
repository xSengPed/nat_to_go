package router

import (
	"nat_backend_go/controllers"
	"nat_backend_go/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {
	dataFeedRepo := repository.NewDataFeedingRepo(db)
	dataFeedController := controllers.InitController(*dataFeedRepo)
	dataFeedGroup := e.Group("/v1")
	dataFeedGroup.POST("/upload", dataFeedController.UploadExcel)
}
