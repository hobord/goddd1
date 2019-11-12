package http

import (
	"net/http"

	"github.com/hobord/goddd1/usecase"

	"github.com/hobord/goddd1/delivery/http/handlers"
)

func MakeRouting(mux *http.ServeMux, entityInteractor *usecase.ExampleInteractor) {

	entityApp := handlers.NewEntityHTTPApp(entityInteractor)

	mux.Handle("/entity/get", entityApp.Get)
	mux.Handle("/entity/getall", entityApp.GetAll)

}
