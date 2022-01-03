package routes

// user will hit when front-end is triggered
import (
	"github.com/YutaUtah/RESTAPIBook/pkg/controllers"
	"github.com/gorilla/mux"
)

// for example, when someone hits /book/ is hit, trigger controllers.CreateBook is executed
// https://pkg.go.dev/github.com/gorilla/mux@v1.8.0#Router.HandleFunc
var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
}
