package controllers

import (
	"nat_backend_go/models"
	"nat_backend_go/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type DataRetriverController struct {
	repository repository.IDataRetriverRepo
}

// Initializer Function that create NewController that return Address of Controller
func NewDataRetriverController(repo repository.DataRetriverRepo) *DataRetriverController {

	return &DataRetriverController{
		repository: repo,
	}

}

// Bound GetCompetitor function in to Controller make function able to call GetCompetitor
func (controller DataRetriverController) GetCompetitors(ctx echo.Context) error {

	data, err := controller.repository.GetData()

	if err != nil {
		return ctx.JSON(http.StatusOK, models.CommonResponse{
			Code:    2000,
			Message: "Unable Get Data",
		})
	}

	return ctx.JSON(http.StatusOK, models.CommonResponse{
		Code:    1000,
		Message: data,
	})
}

func (controller DataRetriverController) GetCompetitorByCid(ctx echo.Context) error {

	cid := ctx.QueryParam("cid")

	competitor, err := controller.repository.GetCompetitorByCid(cid)
	if err != nil {
		return ctx.JSON(http.StatusOK, models.CommonResponse{
			Code:    2000,
			Message: "Unable Get Data",
		})
	}
	return ctx.JSON(http.StatusOK, models.CommonResponse{
		Code: 1000,
		Data: competitor,
	})
}

func (controller DataRetriverController) GetCompetitorsPaginate(ctx echo.Context) error {

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))

	result, rowsCount, err := controller.repository.GetCompetitorPaginate(page, limit)
	if err != nil {
		return ctx.JSON(http.StatusOK, models.CommonResponse{
			Code:    2000,
			Message: "Unable Get Data",
		})
	}
	return ctx.JSON(http.StatusOK, models.PaginationResponse{
		Code:      1000,
		Data:      result,
		RowsCount: int64(rowsCount),
		Page:      int64(page),
		Limit:     int64(limit),
	})
}

func (controller DataRetriverController) GetCompetitorByRegion(ctx echo.Context) error {

	region := ctx.QueryParam("region")

	page, _ := strconv.Atoi(ctx.QueryParam("page"))
	limit, _ := strconv.Atoi(ctx.QueryParam("limit"))
	result, err := controller.repository.GetCompetitorByRegion(page, limit, region)

	if err != nil {
		panic(err)
	}

	return ctx.JSON(http.StatusOK, models.PaginationResponse{
		Code:      1000,
		Data:      result,
		RowsCount: 0,
		Page:      0,
		Limit:     0,
	})
}
