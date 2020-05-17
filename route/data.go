package route

import (
    "buzzer/m/v2/buzzer"
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "log"
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
    //TODO Get Data from DB
}

func PostBuzzer(c *gin.Context){
    var data buzzer.Message
    err := c.BindJSON(&data)
    data.Uuid, _ = uuid.NewRandom()

    if err != nil {
        log.Println(err)
        c.JSON(500, err)
    }

    log.Println(data)
    c.JSON(200, data)
    //TODO Store data in DB
}
