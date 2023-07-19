package controllers

import (
	"errors"
	"fmt"
	"io"
	"nat_backend_go/models"
	"nat_backend_go/repository"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

type DataFeedingController struct {
	repository repository.IDataFeedingRepository
}

func contain(list []string, str string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func InitController(repo repository.DataFeedingRepository) *DataFeedingController {
	return &DataFeedingController{
		repository: repo,
	}
}

func (controller DataFeedingController) UploadExcel(ctx echo.Context) error {

	var subjects []string = []string{"math", "sci", "eng"}
	subject := ctx.QueryParam("subject")

	if contain(subjects, subject) {

		file, err := ctx.FormFile("file")

		if err != nil {
			return errors.New("Cannot Upload")
		}

		src, err := file.Open()

		if err != nil {
			return errors.New("Src Error")
		}

		defer src.Close()

		dst, err := os.Create(file.Filename)

		defer dst.Close()
		_, err = io.Copy(dst, src)

		if err != nil {
			return err
		}

		excelResult, err := excelize.OpenFile(file.Filename)

		if err != nil {
			return err
		}

		competitors := []models.MathCompetitor{}
		sheetName := excelResult.GetSheetName(0)
		rows, err := excelResult.GetRows(sheetName)

		for i := 1; i < len(rows); i++ {

			var competitor models.MathCompetitor

			cell := fmt.Sprintf("A%d", i)
			competitor.Cid, err = excelResult.GetCellValue(sheetName, cell)

			cell = fmt.Sprintf("B%d", i)
			competitor.Name, err = excelResult.GetCellValue(sheetName, cell)

			cell = fmt.Sprintf("C%d", i)
			competitor.Level, err = excelResult.GetCellValue(sheetName, cell)

			cell = fmt.Sprintf("D%d", i)
			competitor.Catergory, err = excelResult.GetCellValue(sheetName, cell)

			cell = fmt.Sprintf("E%d", i)
			competitor.Arena, err = excelResult.GetCellValue(sheetName, cell)

			cell = fmt.Sprintf("F%d", i)
			competitor.School, err = excelResult.GetCellValue(sheetName, cell)

			cell = fmt.Sprintf("G%d", i)
			competitor.Province, err = excelResult.GetCellValue(sheetName, cell)

			cell = fmt.Sprintf("H%d", i)

			competitor.Region, err = excelResult.GetCellValue(sheetName, cell)

			cell = fmt.Sprintf("I%d", i)
			v, _ := excelResult.GetCellValue(sheetName, cell)

			conv, _ := strconv.ParseFloat(v, 32)
			competitor.CalculationSection = float32(conv)

			cell = fmt.Sprintf("J%d", i)
			v, _ = excelResult.GetCellValue(sheetName, cell)
			conv, _ = strconv.ParseFloat(v, 32)
			competitor.ProblemSection = float32(conv)

			cell = fmt.Sprintf("K%d", i)
			v, _ = excelResult.GetCellValue(sheetName, cell)
			conv, _ = strconv.ParseFloat(v, 32)
			competitor.AppliedSection = float32(conv)

			cell = fmt.Sprintf("L%d", i)
			v, _ = excelResult.GetCellValue(sheetName, cell)
			conv, _ = strconv.ParseFloat(v, 32)
			competitor.Score = float32(conv)

			cell = fmt.Sprintf("M%d", i)
			v, _ = excelResult.GetCellValue(sheetName, cell)
			conv, _ = strconv.ParseFloat(v, 32)
			competitor.ProvinceRank = int(conv)

			cell = fmt.Sprintf("N%d", i)
			v, _ = excelResult.GetCellValue(sheetName, cell)
			conv, _ = strconv.ParseFloat(v, 32)
			competitor.RegionalRank = int(conv)

			uuid := uuid.New()
			competitor.Uuid = uuid.String()

			competitors = append(competitors, competitor)

		}

		err = controller.repository.InsertToDB(competitors)

		if err != nil {
			return ctx.JSON(http.StatusOK, models.CommonResponse{
				Code:    4001,
				Message: "Insert Error",
			})
		}

		return ctx.JSON(http.StatusOK, models.CommonResponse{
			Code:    1000,
			Message: "Upload Complete",
			Data: models.UploadResponse{
				RowsCount: len(competitors),
			},
		})
	} else {
		return ctx.JSON(http.StatusOK, models.CommonResponse{
			Code:    5000,
			Message: "Invalid Subject",
		})
	}

}
