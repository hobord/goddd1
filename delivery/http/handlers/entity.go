package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/hobord/goddd1/domain"
	"github.com/hobord/goddd1/usecase"

	"github.com/hobord/goddd1/delivery/http/dto"
)

// EntityHTTPApp is handle the entity related http request responses
type EntityHTTPApp struct {
	entityInteractor *usecase.ExampleInteractor
}

// NewEntityHTTPApp create a new http handler app to entity
func NewEntityHTTPApp(entityInteractor *usecase.ExampleInteractor) *EntityHTTPApp {
	return &EntityHTTPApp{
		entityInteractor: entityInteractor,
	}
}

// Get return entity by id
func (app *EntityHTTPApp) Get(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/entity/")

	entity := app.entityInteractor.Get(string(id))

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
func (app *EntityHTTPApp) GetAll(w http.ResponseWriter, r *http.Request) {}

// Create is update to persistent the entity
func (app *EntityHTTPApp) Create(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var createDTO dto.EntityCreateRequest
	err := decoder.Decode(&createDTO)
	if err != nil {
		return
	}

	entity, err := domain.NewEntity(createDTO.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = app.entityInteractor.Save(entity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Update is update to persistent the entity
func (app *EntityHTTPApp) Update(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var updateDTO dto.EntityUpdateRequest
	err := decoder.Decode(&updateDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	entity, err := app.entityInteractor.Get(context.TODO(), updateDTO.ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entity.Title = updateDTO.Title

	err = app.entityInteractor.Save(entity)
	if err != nil {
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

// Delete entity
func (app *EntityHTTPApp) Delete(w http.ResponseWriter, r *http.Request) {}

const MessageContextKey = "message"

func (app *EntityHTTPApp) AddMessageMiddleware(message string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), MessageContextKey, message)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
