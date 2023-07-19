package routes

import (
	"todo/api/controllers"

	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	route := mux.NewRouter()

	route.HandleFunc("/", controllers.Show)

	route.HandleFunc("/", controllers.Create).Methods("POST")

	route.HandleFunc("/delete/{id}", controllers.Delete)

	route.HandleFunc("/complete/{id}", controllers.Complete)

	return route
}
