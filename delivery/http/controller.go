package http

import (
	"github.com/gorilla/mux"

	"github.com/hobord/goddd1/usecase"

	"github.com/hobord/goddd1/delivery/http/handlers"
)

// MakeRouting is add handler functions to mux router
func MakeRouting(router *mux.Router, entityInteractor *usecase.ExampleInteractor) {
	entityApp := handlers.NewEntityHTTPApp(entityInteractor)

	router.HandleFunc("/entity", entityApp.Create).Methods("POST")
	router.HandleFunc("/entity/{id}", entityApp.Get)
	router.HandleFunc("/entity", entityApp.GetAll).Methods("GET")
	router.HandleFunc("/entity", entityApp.Update).Methods("PUT")
	router.HandleFunc("/entity/{id}", entityApp.Delete).Methods("DELETE")
}
