package router

import (
	"nat_backend_go/controllers"
	"nat_backend_go/repository"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Init(e *echo.Echo, db *gorm.DB) {

	endPoint := e.Group("/api/v1")
	dataFeedRepo := repository.NewDataFeedingRepo(db)
	dataFeedController := controllers.InitController(*dataFeedRepo)
	dataFeedGroup := endPoint.Group("/feed")
	dataFeedGroup.POST("/upload", dataFeedController.UploadExcel)
	dataRetRepo := repository.NewDataRetriverRepo(db)
	dataRetController := controllers.NewDataRetriverController(*dataRetRepo)
	dataRetGroup := endPoint
	dataRetGroup.GET("/get_competitor", dataRetController.GetCompetitors)
	dataRetGroup.GET("/get_competitor_by_cid", dataRetController.GetCompetitorByCid)
	dataRetGroup.GET("/get_competitor_paginate", dataRetController.GetCompetitorsPaginate)
	dataRetGroup.GET("/get_competitor_by_region", dataRetController.GetCompetitorByRegion)

}
