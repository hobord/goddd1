package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hobord/goddd1/domain"
	"github.com/hobord/goddd1/usecase"

	"github.com/hobord/goddd1/delivery/http/dto"
)

// TODO: make custom http errors

// EntityHTTPApp is handle the entity related http request responses
type EntityHTTPApp struct {
	entityInteractor usecase.ExampleInteractorInterface
}

// NewEntityHTTPApp create a new http handler app to entity
func NewEntityHTTPApp(entityInteractor usecase.ExampleInteractorInterface) *EntityHTTPApp {
	return &EntityHTTPApp{
		entityInteractor: entityInteractor,
	}
}

// GetByID return entity by id
func (app *EntityHTTPApp) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	entity, err := app.entityInteractor.GetByID(r.Context(), string(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		err = errors.New("No resource found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entityDTO := &dto.EntityResponse{
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
func (app *EntityHTTPApp) GetAll(w http.ResponseWriter, r *http.Request) {
	entities, err := app.entityInteractor.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entities == nil || len(entities) == 0 {
		err = errors.New("No resource found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entityDTOs := make([]dto.EntityResponse, 1)
	for _, entity := range entities {
		entityDTO := &dto.EntityResponse{
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
func (app *EntityHTTPApp) Create(w http.ResponseWriter, r *http.Request) {
	// Decode the request DTO.
	decoder := json.NewDecoder(r.Body)
	var createDTO dto.EntityCreateRequest
	err := decoder.Decode(&createDTO)
	if err != nil {
		return
	}

	// Create a new entity.
	entity, err := domain.NewEntity(createDTO.Title)
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
	entityDTO := &dto.EntityResponse{
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
func (app *EntityHTTPApp) Update(w http.ResponseWriter, r *http.Request) {
	// Decode the request DTO.
	decoder := json.NewDecoder(r.Body)
	var updateDTO dto.EntityUpdateRequest
	err := decoder.Decode(&updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Load the original entity.
	entity, err := app.entityInteractor.GetByID(r.Context(), updateDTO.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		err = errors.New("No resource found")
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
	entityDTO := &dto.EntityResponse{
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
func (app *EntityHTTPApp) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	entity, err := app.entityInteractor.GetByID(r.Context(), string(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
func (app *EntityHTTPApp) AddMessageMiddleware(message string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), MessageContextKey, message)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
