package main

import (
	"buzzer/m/v2/repository"
	"buzzer/m/v2/route"
	"github.com/gin-gonic/gin"
	"log"
)

var redisRepository repository.Redis

func main() {

	redisRepository := new(repository.Redis)
	redisRepository.Init()
	ping, err := redisRepository.Ping()

	if err != nil || ping != "PONG" {
		log.Fatal(err)
	}

	r := gin.Default()
	r.Static("/static", "./static")
	r.LoadHTMLGlob("template/*")
	r.GET("/", route.GetEntry)
	rData := r.Group("/data")
	{
		rData.GET("/buzzer", route.GetAllBuzzers)
		rData.POST("/buzzer", route.PostBuzzer)
	}

	r.GET("/admin", route.GetAdmin)
	r.GET("/monitoring", route.GetMonitoring)

	r.Run(":3000")
}
