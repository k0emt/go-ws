package bite

import (
	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Echo(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": c.Param("text"),
	})
}
