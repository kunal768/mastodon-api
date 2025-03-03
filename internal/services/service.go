package services

import (
	"context"
	"mastadon-api/internal/client/mastadon"
	"mastadon-api/internal/models"
)

type Service interface {
	CreateNewPost(ctx context.Context, Content mastadon.CreatePost) (models.CreatePostResponse, error)
	DeletePost(ctx context.Context, Req mastadon.DeletePost) (bool, error)
	FetchPosts(ctx context.Context, Req mastadon.FetchPosts) ([]mastadon.StatusHistory, error)
}

type svc struct {
	mastSvc mastadon.Service
}

func NewHTTPService(mastSvc mastadon.Service) Service {
	return &svc{mastSvc: mastSvc}
}

func (s svc) CreateNewPost(ctx context.Context, Content mastadon.CreatePost) (models.CreatePostResponse, error) {
	res, err := s.mastSvc.CreateNewPost(Content)
	var ans models.CreatePostResponse
	if err != nil {
		ans.Success = false
		ans.Message = err.Error()
		return ans, err
	}
	return models.CreatePostResponse{Success: true, Message: "success", StatusId: res.StatusId}, nil
}

func (s svc) DeletePost(ctx context.Context, Req mastadon.DeletePost) (bool, error) {
	res, err := s.mastSvc.DeletePost(Req)
	if err != nil {
		return false, err
	}
	return res.Success, nil
}

func (s svc) FetchPosts(ctx context.Context, Req mastadon.FetchPosts) ([]mastadon.StatusHistory, error) {
	res, err := s.mastSvc.FetchPosts(Req)
	if err != nil {
		return nil, err
	}
	return res.History, nil
}
