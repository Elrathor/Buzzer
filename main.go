package main

import (
	"buzzer/m/v2/route"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("template/*")
	r.GET("/", route.GetEntry)
	rData := r.Group("/data")
	{
		rData.GET("/buzzer", route.GetAllBuzzers)
	}

	r.GET("/admin", route.GetAdmin)
	r.GET("/monitoring", route.GetMonitoring)

	r.Run(":3000")
}
