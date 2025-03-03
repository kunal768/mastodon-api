package models

type CreatePostResponse struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	StatusId string `json:"status_id"`
}
