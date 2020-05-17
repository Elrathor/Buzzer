package route

import (
	"buzzer/m/v2/buzzer"
	"buzzer/m/v2/repository"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"sort"
)

var redisRepository repository.Redis

func GetAllBuzzers(c *gin.Context) {
	lazyInitializeDB()
	buzzers, err := redisRepository.GetAllMessages()

	if err != nil {
		c.JSON(500, err)
	}

	//Order from Old to New

	sort.SliceStable(buzzers[:], func(i, j int) bool {
		return buzzers[j].BuzzTime.After(buzzers[i].BuzzTime)
	})

	c.JSON(200, buzzers)
}

func PostBuzzer(c *gin.Context) {
	lazyInitializeDB()
	var data buzzer.Message
	err := c.BindJSON(&data)
	data.Uuid, _ = uuid.NewRandom()

	if err != nil {
		log.Println(err)
		c.JSON(500, err)
	}

	data, err = redisRepository.SetMessage(data)

	if err != nil {
		log.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, data)
}

func DeleteBuzzers(c *gin.Context) {
	lazyInitializeDB()
	result, err := redisRepository.DeleteAllMessages()

	if err != nil {
		log.Println(err)
		c.JSON(500, err)
	}

	c.JSON(200, gin.H{
		"numberOfElements": result,
	})
}

func lazyInitializeDB() {
	if !redisRepository.GetIsInitialized() {
		redisRepository.Init()
	}
}
