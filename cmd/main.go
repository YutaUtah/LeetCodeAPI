package main

import (
	"log"
	"net/http"

	"github.com/YutaUtah/RESTAPIBook/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/zinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}
