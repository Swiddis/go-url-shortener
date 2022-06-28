package main

import (
	"fmt"

	"github.com/Swiddis/go-url-shortener/handler"
	"github.com/Swiddis/go-url-shortener/store"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/health/ready", func(c *gin.Context) {
		handler.IsReady(c)
	})
	r.GET("/health/live", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "LIVE",
		})
	})

	r.POST("/", func(c *gin.Context) {
		handler.CreateShortUrl(c)
	})
	r.GET("/:shortUrl", func(c *gin.Context) {
		handler.HandleShortUrlRedirect(c)
	})

	store.InitializeStore()

	err := r.Run(":9808")
	if err != nil {
		panic(fmt.Sprintf("Failed to start the web server - Error: %v", err))
	}
}
