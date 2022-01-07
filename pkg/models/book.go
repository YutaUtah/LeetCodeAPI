package models

import (
	"github.com/YutaUtah/RESTAPIBook/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Problem struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	// automigrate empty book
	db.AutoMigrate(&Problem{})
}

// functions to communicate with database

func (b *Problem) CreateProblem() *Problem {
	// gorm function
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllProblems() []Problem {
	// slice of Book struct
	var Problems []Problem
	db.Find(&Problems)
	return Problems
}

func GetProblemById(Id int64) (*Problem, *gorm.DB) {
	var getProblem Problem
	db := db.Where("ID=?", Id).Find(&getProblem)
	return &getProblem, db
}

func DeleteProblem(ID int64) Problem {
	var problem Problem
	db.Where("ID=?", ID).Delete(problem)
	return problem
}
