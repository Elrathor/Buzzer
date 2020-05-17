package route

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetAdmin (c *gin.Context) {
    c.HTML(http.StatusOK, "admin.tmpl", gin.H{
        "title": "This is a test",
    })
}