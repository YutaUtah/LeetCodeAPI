package main

import (
	"log"
	"net/http"

	"github.com/YutaUtah/RESTAPIBook/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// initialize router
	r := mux.NewRouter()
	// function with multiple routing options
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	// check this later help us create server in the first place
	// if there is error print
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}
