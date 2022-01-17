package config

// this file returns variable db

import (
	"fmt"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// App config
type App struct {
	Router     *mux.Router
	PathPrefix string
}

func (a *App) Getrouter() *mux.Router {
	if a.Router == nil {
		a.Router = mux.NewRouter().StrictSlash(true)
	}
	return a.Router
}

func (a *App) GetSubrouter(PathPrefix string) *mux.Router {
	return a.Getrouter().PathPrefix(a.PathPrefix).Subrouter()
}

// db config
var (
	db *gorm.DB
)

func argInit() string {
	dsn := fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}

func Connect() {
	// this enables you to connect with database: username:password:tablename
	dsn := argInit()
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
