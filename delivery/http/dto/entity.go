package dto

// EntityCreateRequest is request DTO for create entity with http
type EntityCreateRequest struct {
	Title string `json:"title"`
}

// EntityUpdateRequest is request DTO for update entity with http
type EntityUpdateRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// EntityResponse is response DTO for entity with http
type EntityResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// EntitiesResponse is response DTO for multiple entities
type EntitiesResponse struct {
	Entities []EntityResponse `json:"entities"`
}
