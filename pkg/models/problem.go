package models

import (
	"time"

	"github.com/YutaUtah/LeetCodeAPI/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Problem struct {
	gorm.Model
	Number     int    `gorm:"" json:"number"`
	Title      string `json:"title"`
	Difficulty string `json:"difficulty"`
	FinishedDT *Date  `json:"date"`
}

type Date struct {
	time.Time
}

// custom UnmarshalJSON to specify the date format for FinishedDT
func (d *Date) UnmarshalJSON(b []byte) error {
	v, err := time.Parse(`"2006-01-02"`, string(b))
	if err != nil {
		return err
	}
	d.Time = v
	return nil
}

func init() {
	config.Connect()
	db = config.GetDB()
	// automigrate empty problem
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
	// slice of Problem struct
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
