package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/hobord/goddd1/domain"
	"github.com/hobord/goddd1/usecase/mocks"
	"github.com/icrowley/fake"
	mock "github.com/stretchr/testify/mock"
)

func TestGetByID(t *testing.T) {

	// Create a mock uses case interactor and mock the results
	mockUsecase := &mocks.ExampleInteractorInterface{}

	fakeID := fake.Sentence()
	fakeTitle := fake.Sentence()
	usecaseReturnEntity := &domain.Entity{
		ID:    fakeID,
		Title: fakeTitle,
	}
	mockUsecase.On("GetByID", mock.Anything, mock.Anything).Return(usecaseReturnEntity, nil)

	// Create a test HTTPApp with moc usecase
	app := NewEntityHTTPApp(mockUsecase)

	// Create a test request
	req, err := http.NewRequest("GET", "/entity/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()

	// Create a router and assign our handler func
	router := mux.NewRouter()
	router.HandleFunc("/entity/{id}", app.GetByID)

	// Make a request into the router
	router.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := fmt.Sprintf(`{"id":"%s","title":"%s"}`, fakeID, fakeTitle)
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// TODO: implement all handler func tests
