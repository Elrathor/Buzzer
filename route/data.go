package route

import (
    "buzzer/m/v2/buzzer"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "time"
)

func GetData (c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "pong",
    })
}

func GetAllBuzzers (c *gin.Context) {
    uuidVal, _ := uuid.NewRandom()
    buzzerOne := buzzer.Message{
        Uuid: uuidVal,
        TeamName:   "Elephant",
        PlayerName: "Tobias",
        BuzzTime:   time.Now(),
    }
    buzzers := [...]buzzer.Message{buzzerOne, buzzerOne}

    c.JSON(200, buzzers)
}
