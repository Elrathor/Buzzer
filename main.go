package main

import (
    "buzzer/m/v2/route"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("template/*")
    r.GET("/", route.GetEntry)
    r.GET("/data", route.GetData)
    r.GET("/admin", route.GetAdmin)
    r.GET("/monitoring", route.GetMonitoring)

    r.Run(":3000")
}