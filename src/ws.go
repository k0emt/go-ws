package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	services "web-spike/services"
)

func main() {
	fmt.Println("Starting web service")

	router := gin.Default()
	router.Static("/assets", "../assets")
	router.Static("images", "../assets/images") // more direct path to images
	router.StaticFS("/static", http.Dir("../static"))
	router.StaticFile("/favicon.ico", "../resources/favicon.ico")

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/articles", services.GetArticles)
	router.GET("/articles/:id", services.GetArticleByID)
	router.POST("/articles", services.PostArticles)

	router.GET("/UTC", services.GetUtcTime)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
