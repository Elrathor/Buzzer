package route

import "github.com/gin-gonic/gin"

func GetAdmin (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}