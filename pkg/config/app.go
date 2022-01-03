package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// todo: this can be written in a different way
var (
	db *gorm.DB
)

func Connect() {
	// username:password:tablename
	d, err := gorm.Open("mysql", "root:Yutagenta55@(localhost:3306)/bookinfo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
