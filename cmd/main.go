package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/YutaUtah/LeetCodeAPI/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// initialize router
	var addr = flag.String("addr", ":9000", "endpoint address")
	r := mux.NewRouter().StrictSlash(true)
	// function with multiple routing options
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Println("Starting web server on", *addr)
	log.Fatal(http.ListenAndServe("localhost:9000", r))
}
