package dto

// FooEntityCreateRequest is request DTO for create entity with http
type FooEntityCreateRequest struct {
	Title string `json:"title"`
}

// FooEntityUpdateRequest is request DTO for update entity with http
type FooEntityUpdateRequest struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// FooEntityResponse is response DTO for entity with http
type FooEntityResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

// FooEntitiesResponse is response DTO for multiple entities
type FooEntitiesResponse struct {
	Entities []FooEntityResponse `json:"entities"`
}
