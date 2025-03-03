package apis

import (
	"net/http"
	"os"
	"strconv"

	"mastadon-api/internal/client/mastadon"
	"mastadon-api/internal/services"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

// HTTP handlers for Mastodon API endpoints
func createPostHandler(service services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req mastadon.CreatePost
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "invalid request format",
				"details": err.Error(),
			})
			return
		}

		// Validate required fields
		if req.Status == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "missing required field: status",
			})
			return
		}

		// Call service layer
		success, err := service.CreateNewPost(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to create post",
				"details": err.Error(),
			})
			return
		}

		// Return success response
		c.JSON(http.StatusCreated, gin.H{
			"success":   success,
			"message":   "Post created successfully",
			"status_id": success.StatusId,
		})
	}
}

func deletePostHandler(service services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		_, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "invalid post ID",
				"details": "ID must be a valid integer",
			})
			return
		}

		req := mastadon.DeletePost{Id: idParam}
		success, err := service.DeletePost(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to delete post",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": success,
			"message": "Post deleted successfully",
		})
	}
}

func fetchPostsHandler(service services.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		req := mastadon.FetchPosts{}
		posts, err := service.FetchPosts(c.Request.Context(), req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "failed to fetch posts",
				"details": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": posts,
		})
	}
}

// RegisterHandlers registers all handlers with the router
func RegisterHandlers(router *gin.Engine, service services.Service) {
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("ALLOWED_ORIGIN")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	router.POST("/posts", createPostHandler(service))
	router.DELETE("/posts/:id", deletePostHandler(service))
	router.GET("/posts", fetchPostsHandler(service))
}
