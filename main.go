package main

import (
	"buzzer/m/v2/repository"
	"buzzer/m/v2/route"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	libredis "github.com/go-redis/redis/v7"
	limiter "github.com/ulule/limiter/v3"
	mgin "github.com/ulule/limiter/v3/drivers/middleware/gin"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
	"log"
)

func main() {
	// Check connection to REDIS Database or panic
	redisRepository := new(repository.Redis)
	redisRepository.Init()
	ping, err := redisRepository.Ping()

	if err != nil || ping != "PONG" {
		log.Fatal(err)
	}

	// Initialize rate limiter
	rate, err := limiter.NewRateFromFormatted("120-M")
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a redis client.
	option, err := libredis.ParseURL("redis://redis-master:6379/1")
	if err != nil {
		log.Fatal(err)
		return
	}
	client := libredis.NewClient(option)

	// Create a store with the redis client.
	store, err := sredis.NewStoreWithOptions(client, limiter.StoreOptions{
		Prefix:   "limiter_gin_example",
		MaxRetry: 3,
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	// Create a new rateLimmeter with the limiter instance.
	rateLimeter := mgin.NewMiddleware(limiter.New(store, rate))

	// CORS Configuration
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowMethods = []string{"GET", "POST", "DELETE"}
	corsConfig.AllowOrigins = []string{
		"https://jsdelivr.net",
		"https://bootstrapcdn.com",
		"https://jquery.com",
	}

	r := gin.Default()
	r.ForwardedByClientIP = true
	r.Use(rateLimeter)
	r.Use(cors.New(corsConfig))
	r.Static("/static", "./static")
	r.LoadHTMLGlob("template/*")
	r.GET("/", route.GetEntry)
	rData := r.Group("/data")
	{
		rData.GET("/buzzer", route.GetAllBuzzers)
		rData.POST("/buzzer", route.PostBuzzer)
		rData.DELETE("/buzzer", route.DeleteBuzzers)
	}

	rAdmin := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "bee",
	}))

	rAdmin.GET("/", route.GetAdmin)
	r.GET("/monitoring", route.GetMonitoring)

	err = r.Run(":3000")

	if err != nil {
		log.Panic(err)
	}
}
