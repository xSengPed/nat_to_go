package repository

import (
	"fmt"

	"gorm.io/gorm"
)

type DataRetriverRepo struct {
	db *gorm.DB
}
type IDataRetriverRepo interface {
	GetData() (string, error)
}

// create repo with db gorm which argument is gorm db as instance and return address of struct
func NewDataRetriverRepo(db *gorm.DB) *DataRetriverRepo {
	return &DataRetriverRepo{
		db: db,
	}
}

// receiver function bond GetData() in to DataRetriver Struct it behavior same like class in oop
func (repo DataRetriverRepo) GetData() (string, error) {
	fmt.Println("Repo Get Data")
	return "OK", nil
}
