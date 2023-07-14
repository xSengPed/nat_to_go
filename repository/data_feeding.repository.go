package repository

import (
	"nat_backend_go/models"

	"gorm.io/gorm"
)

type DataFeedingRepository struct {
	db *gorm.DB
}

type IDataFeedingRepository interface {
	InsertToDB(competitors []models.MathCompetitor) error
}

func NewDataFeedingRepo(db *gorm.DB) *DataFeedingRepository {
	// Auto Create Table
	db.AutoMigrate(&models.MathCompetitor{})
	return &DataFeedingRepository{
		db: db,
	}
}

func (repo DataFeedingRepository) InsertToDB(competitors []models.MathCompetitor) error {

	for idx, v := range competitors {
		if idx != 0 {
			err := repo.db.Create(v).Error

			if err != nil {
				return err
			}
		}
	}

	return nil

}
