package mastadon

import "time"

type CreatePost struct {
	Status string `json:"status"`
}

type DeletePost struct {
	Id string `json:"id"`
}

type FetchPosts struct {
}

type Response struct {
	Success  bool   `json:"success"`
	StatusId string `json:"status_id"`
}

type FetchPostsResponse struct {
	History []StatusHistory `json:"history"`
}

type StatusHistory struct {
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	StatusId  string    `json:"status_id"`
}
