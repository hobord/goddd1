package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hobord/goddd1/delivery/http/dto"
	entities "github.com/hobord/goddd1/domain/entity"
	"github.com/hobord/goddd1/usecase"
)

// TODO: make custom http errors

// FooEntityRestHTTPModule is handle the entity related http request responses
type FooEntityRestHTTPModule struct {
	entityInteractor usecase.ExampleInteractorInterface
}

// NewFooEntityRestHTTPModule create a new http handler app to entity
func NewFooEntityRestHTTPModule(entityInteractor usecase.ExampleInteractorInterface) *FooEntityRestHTTPModule {
	return &FooEntityRestHTTPModule{
		entityInteractor: entityInteractor,
	}
}

// GetByID return entity by id
func (app *FooEntityRestHTTPModule) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	entity, err := app.entityInteractor.GetByID(r.Context(), string(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		err = errors.New("No resource found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	entityDTO := &dto.FooEntityResponse{
		ID:    entity.ID,
		Title: entity.Title,
	}
	js, err := json.Marshal(entityDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// GetAll return all entities
func (app *FooEntityRestHTTPModule) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := app.entityInteractor.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if res == nil || len(res) == 0 {
		err = errors.New("No resource found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	entityDTOs := make([]dto.FooEntityResponse, 1)
	for _, entity := range res {
		entityDTO := &dto.FooEntityResponse{
			ID:    entity.ID,
			Title: entity.Title,
		}
		entityDTOs = append(entityDTOs, *entityDTO)
	}

	// Convert to json
	js, err := json.Marshal(entityDTOs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back to response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Create is update to persistent the entity
func (app *FooEntityRestHTTPModule) Create(w http.ResponseWriter, r *http.Request) {
	// Decode the request DTO.
	decoder := json.NewDecoder(r.Body)
	var createDTO dto.FooEntityCreateRequest
	err := decoder.Decode(&createDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new entity.
	entity, err := entities.NewFooEntity(createDTO.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Save the new entity.
	err = app.entityInteractor.Save(r.Context(), entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a new response DTO.
	entityDTO := &dto.FooEntityResponse{
		ID:    entity.ID,
		Title: entity.Title,
	}
	// Convert to json
	js, err := json.Marshal(entityDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back to response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Update is update to persistent the entity.
func (app *FooEntityRestHTTPModule) Update(w http.ResponseWriter, r *http.Request) {
	// Decode the request DTO.
	decoder := json.NewDecoder(r.Body)
	var updateDTO dto.FooEntityUpdateRequest
	err := decoder.Decode(&updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Load the original entity. TODO: move to usecase?
	entity, err := app.entityInteractor.GetByID(r.Context(), updateDTO.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		err = errors.New("No resource found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Update the entity properties.
	entity.Title = updateDTO.Title

	// save the entity
	err = app.entityInteractor.Save(r.Context(), entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a response DTO.
	entityDTO := &dto.FooEntityResponse{
		ID:    entity.ID,
		Title: entity.Title,
	}
	// Convert to json
	js, err := json.Marshal(entityDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send back to response.
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// Delete entity
func (app *FooEntityRestHTTPModule) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	entity, err := app.entityInteractor.GetByID(r.Context(), string(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	err = app.entityInteractor.Delete(r.Context(), entity.ID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// MessageContextKey is an unique context key for example message context
const MessageContextKey = "message"

// AddMessageMiddleware is an example message middleware
func (app *FooEntityRestHTTPModule) AddMessageMiddleware(message string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), MessageContextKey, message)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
