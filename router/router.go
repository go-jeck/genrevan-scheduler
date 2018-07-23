package router

import (
	"github.com/go-squads/genrevan-scheduler/controller"
	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/lxc", controller.GetLXCs).Methods("GET")

	return router
}
