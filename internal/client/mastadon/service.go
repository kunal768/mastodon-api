package mastadon

import (
	"context"
	"log"

	"github.com/mattn/go-mastodon"
)

type Service interface {
	CreateNewPost(request CreatePost) (Response, error)
	DeletePost(request DeletePost) (Response, error)
	FetchPosts(request FetchPosts) (FetchPostsResponse, error)
}

type svc struct {
	userId string
	client *mastodon.Client
}

func NewService(client *mastodon.Client, userId string) Service {
	return &svc{
		userId: userId,
		client: client,
	}
}

func (s svc) CreateNewPost(request CreatePost) (Response, error) {
	toot := mastodon.Toot{
		Status:     request.Status,
		Visibility: "public",
	}
	status, err := s.client.PostStatus(context.Background(), &toot)
	var res Response
	if err != nil {
		log.Fatalf("%#v\n", err)
		res.Success = false
		return res, err
	}
	res.Success = true
	res.StatusId = string(status.ID)
	return res, nil
}

func (s svc) DeletePost(request DeletePost) (Response, error) {
	err := s.client.DeleteStatus(context.Background(), mastodon.ID(request.Id))
	if err != nil {
		log.Fatalf("%#v\n", err)
		return Response{Success: false}, err
	}
	return Response{Success: true}, nil
}

func (s svc) FetchPosts(request FetchPosts) (FetchPostsResponse, error) {
	pagination := mastodon.Pagination{
		Limit: 10,
	}
	res, err := s.client.GetAccountStatuses(context.Background(), mastodon.ID(s.userId), &pagination)
	var ans FetchPostsResponse
	if err != nil {
		log.Fatalf("%#v\n", err)
		return ans, err
	}
	ans.History = make([]StatusHistory, 0)
	for _, sh := range res {
		converted := StatusHistory{
			Content:   sh.Content,
			CreatedAt: sh.CreatedAt,
		}
		ans.History = append(ans.History, converted)
	}
	return ans, nil
}
