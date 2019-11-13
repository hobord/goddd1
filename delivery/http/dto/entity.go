package dto

type EntityCreateRequest struct {
	Title string `json:"title"`
}

type EntityUpdateRequest struct {
	ID    string `json:"ID"`
	Title string `json:"title"`
}

type EntityResponse struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}
