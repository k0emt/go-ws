package articles

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// article represents data about a written article
type article struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	WordCount int    `json:"wordCount"`
}

// seed data
var articles = []article{
	{ID: "1", Title: "Blue Train", Author: "John Coltrane", WordCount: 5699},
	{ID: "2", Title: "Jedi mind tricks", Author: "Gerry Mulligan", WordCount: 1798},
	{ID: "3", Title: "Something or other", Author: "Sarah Vaughan", WordCount: 3997},
}

func GetArticles(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, articles)
	c.JSON(http.StatusOK, articles)
}

func PostArticles(c *gin.Context) {
	var newArticle article

	if err := c.BindJSON(&newArticle); err != nil {
		return
	}

	articles = append(articles, newArticle)
	c.IndentedJSON(http.StatusCreated, newArticle)
}

func GetArticleByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range articles {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}
