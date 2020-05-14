package route

import "github.com/gin-gonic/gin"

func GetEntry (c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}