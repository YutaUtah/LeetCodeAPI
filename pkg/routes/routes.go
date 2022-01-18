package routes

// user will hit when front-end is triggered
import (
	"github.com/YutaUtah/LeetCodeAPI/pkg/controllers"
	"github.com/gorilla/mux"
)

// for example, when someone hits /book/ is hit, trigger controllers.CreateBook is executed
// https://pkg.go.dev/github.com/gorilla/mux@v1.8.0#Router.HandleFunc
var RegisterLeetCodeRoutes = func(router *mux.Router) {
	router.HandleFunc("/problems/", controllers.CreateProblem).Methods("POST")
	router.HandleFunc("/problems/", controllers.GetProblem).Methods("GET")
	router.HandleFunc("/problems/{Id}", controllers.GetProblemById).Methods("GET")
	router.HandleFunc("/problems/{Id}", controllers.UpdateProblem).Methods("PUT")
	router.HandleFunc("/problems/{Id}", controllers.DeleteProblem).Methods("DELETE")
	router.HandleFunc("/about", controllers.DisplayProblems)
}
