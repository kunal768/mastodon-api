package main

import (
	"mastadon-api/internal/apis"
	"mastadon-api/internal/client/mastadon"
	"mastadon-api/internal/services"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/mattn/go-mastodon"
)

func initMastadonClient() *mastodon.Client {
	config := &mastodon.Config{
		Server:       os.Getenv("MASTADON_SERVER"),
		ClientID:     os.Getenv("CLIENT_KEY"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
	}
	return mastodon.NewClient(config)
}

func initMastadonService(client *mastodon.Client) mastadon.Service {
	userId := os.Getenv("USER_ID")
	return mastadon.NewService(client, userId)
}

func initHttpClient(mastSvc mastadon.Service) services.Service {
	return services.NewHTTPService(mastSvc)

}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	mastadonClient := initMastadonClient()
	mastadonService := initMastadonService(mastadonClient)
	httpSvc := initHttpClient(mastadonService)

	// Create router and register handlers
	router := gin.Default()
	apis.RegisterHandlers(router, httpSvc)

	// Start server
	router.Run(":8080")

}
