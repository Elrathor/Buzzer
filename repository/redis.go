package repository

import (
    "buzzer/m/v2/buzzer"
    "github.com/go-redis/redis"
    "github.com/google/uuid"
)

type Redis struct {
    client *redis.Client
    isInitialized bool
}

func (r *Redis)Init(){
    r.client := redis.NewClient(&redis.Options{
        Addr:     "redis",
        Password: "",
        DB:       0,
    })
}

func(r Redis)GetMessage(id uuid.UUID)(message buzzer.Message){
    panic("Not yet implemented")
}

func(r Redis)SetMessage(message buzzer.Message)(id uuid.UUID){
    panic("Not yet implemented")
}

func(r Redis)GetAllMessages()(messages []buzzer.Message){
    panic("Not yet implemented")
}

func(r Redis)DeleteMessage(id uuid.UUID)(wasDeleted bool){
    panic("Not yet implemented")
}

func(r Redis)DeleteAllMessages()(numberOfMessages int){
    panic("Not yet implemented")
}
