package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/YutaUtah/LeetCodeAPI/pkg/config"
	"github.com/YutaUtah/LeetCodeAPI/pkg/routes"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// initialize router
	var addr = flag.String("addr", ":9000", "endpoint address")

	app := config.App{}
	v1 := app.GetSubrouter("/api/v1")

	// function with multiple routing options
	routes.RegisterLeetCodeRoutes(v1)
	http.Handle("/", v1)
	log.Println("Starting web server on", *addr)
	log.Fatal(http.ListenAndServe("localhost:9000", v1))
}
