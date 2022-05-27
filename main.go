package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/maximovd/url-shortener/handler"
	"github.com/maximovd/url-shortener/store"
)

func main() {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "Hey GO URL Shortener !",
		})
	})

	router.POST("/shorter/url/create", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})

	router.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandlerShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := router.Run(":9008")
	if err != nil {
		panic(fmt.Sprintf("Failed to start ther web server - Error details: %v", err))
	}
}
