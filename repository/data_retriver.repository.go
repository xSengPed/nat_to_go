package repository

import (
	"nat_backend_go/models"

	"gorm.io/gorm"
)

type DataRetriverRepo struct {
	db *gorm.DB
}
type IDataRetriverRepo interface {
	GetData() (string, error)
	GetCompetitorByCid(cid string) (models.MathCompetitor, error)
	GetCompetitorPaginate(page int, limit int) ([]models.MathCompetitor, int, error)
}

// create repo with db gorm which argument is gorm db as instance and return address of struct
func NewDataRetriverRepo(db *gorm.DB) *DataRetriverRepo {
	return &DataRetriverRepo{
		db: db,
	}
}

// receiver function bond GetData() in to DataRetriver Struct it behavior same like class in oop
func (repo DataRetriverRepo) GetData() (string, error) {

	return "OK", nil
}

func (repo DataRetriverRepo) GetCompetitorByCid(cid string) (models.MathCompetitor, error) {

	user := models.MathCompetitor{}
	err := repo.db.First(&user, "cid = ?", cid).Error
	if err != nil {
		return models.MathCompetitor{}, err
	}
	return user, nil
}

func (repo DataRetriverRepo) GetCompetitorPaginate(page int, limit int) ([]models.MathCompetitor, int, error) {
	var competitors []models.MathCompetitor
	offset := (page - 1) * limit

	result := repo.db.Find(&competitors)
	rowsCount := result.RowsAffected

	err := repo.db.Offset(offset).Limit(limit).Find(&competitors).Error
	if err != nil {
		return competitors, 0, err
	}
	return competitors, int(rowsCount), nil

}
