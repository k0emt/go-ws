package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	articles "go-ws/articles"
	bite "go-ws/bite"
	time "go-ws/time"
)

func main() {
	fmt.Println("Starting web service")

	router := gin.Default()
	router.Static("/assets", "../assets")
	router.Static("images", "../assets/images") // more direct path to images
	router.StaticFS("/static", http.Dir("../static"))
	router.StaticFile("/favicon.ico", "../resources/favicon.ico")

	router.GET("/echo/:text", bite.Echo)
	router.GET("/ping", bite.Pong)

	router.GET("/articles", articles.GetArticles)
	router.GET("/articles/:id", articles.GetArticleByID)
	router.POST("/articles", articles.PostArticles)

	router.GET("/UTC", time.GetUtcTime)

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
