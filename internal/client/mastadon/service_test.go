package mastadon

import (
	// "context"
	"testing"

	"github.com/mattn/go-mastodon"
	"github.com/stretchr/testify/assert"
)

var createdPostID string // Global variable to store the created post ID

func TestCreateNewPost(t *testing.T) {
	// Setup Mastodon client (real client, not mock)
	client := mastodon.NewClient(&mastodon.Config{
		Server:      "https://mastodon.social",                     // Replace with your Mastodon instance URL
		AccessToken: "znpqyq7wd5qUe765g4bJbWfQxbpqu2y1vrUneFOtLE0", // Replace with your actual access token
	})

	// Create service
	service := NewService(client, "kenilhv@mastodon.social") // Replace with your user ID

	// Create a new post
	request := CreatePost{Status: "Hello, Mastodon!"}
	response, err := service.CreateNewPost(request)

	assert.NoError(t, err)
	assert.True(t, response.Success)
	assert.NotEmpty(t, response.StatusId) // Assert the response has a StatusId

	// Store the created post ID for use in delete test
	createdPostID = response.StatusId
}

func TestDeletePost(t *testing.T) {
	// Setup Mastodon client (real client, not mock)
	client := mastodon.NewClient(&mastodon.Config{
		Server:      "https://mastodon.social",                     // Replace with your Mastodon instance URL
		AccessToken: "znpqyq7wd5qUe765g4bJbWfQxbpqu2y1vrUneFOtLE0", // Replace with your actual access token
	})

	// Create service
	service := NewService(client, "kenilhv@mastodon.social") // Replace with your user ID

	// Ensure that a post was created in the previous test
	if createdPostID == "" {
		t.Fatal("No post was created. Cannot delete.")
	}

	// Delete the post that was created in the previous test
	deleteRequest := DeletePost{Id: createdPostID}
	deleteResponse, err := service.DeletePost(deleteRequest)

	assert.NoError(t, err)
	assert.True(t, deleteResponse.Success) // Assert deletion was successful
}

func TestFetchPosts(t *testing.T) {
	client := mastodon.NewClient(&mastodon.Config{
		Server:      "https://mastodon.social",
		AccessToken: "znpqyq7wd5qUe765g4bJbWfQxbpqu2y1vrUneFOtLE0n",
	})

	// Create service
	service := NewService(client, "114095510376042790")

	// Fetch posts
	request := FetchPosts{}
	response, err := service.FetchPosts(request)

	assert.NoError(t, err)
	assert.NotEmpty(t, response.History)
	assert.Equal(t, "<p>Hello, Mastodon!</p>", response.History[0].Content)
}
