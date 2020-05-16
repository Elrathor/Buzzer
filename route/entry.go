package route

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func GetEntry (c *gin.Context) {
    c.HTML(http.StatusOK, "index.tmpl", gin.H{
        "title": "This is a test",
    })
}