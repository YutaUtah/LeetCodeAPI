package routes

// user will hit when front-end is triggered
import (
	"github.com/YutaUtah/RESTAPIBook/pkg/controllers"
	"github.com/gorilla/mux"
)

// for example, when someone hits /book/ is hit, trigger controllers.CreateBook is executed
// https://pkg.go.dev/github.com/gorilla/mux@v1.8.0#Router.HandleFunc
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/problem/", controllers.CreateProblem).Methods("POST")
	router.HandleFunc("/problem/", controllers.GetProblem).Methods("GET")
	router.HandleFunc("/problem/{Id}", controllers.GetProblemById).Methods("GET")
	router.HandleFunc("/problem/{Id}", controllers.UpdateProblem).Methods("PUT")
	router.HandleFunc("/problem/{Id}", controllers.DeleteProblem).Methods("DELETE")
}
