package repository

import (
	"buzzer/m/v2/buzzer"
	"buzzer/m/v2/util"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"time"
)

type Redis struct {
	client        *redis.Client
	isInitialized bool
}

func (r *Redis) Init() {
	r.client = redis.NewClient(&redis.Options{
		Addr:     "redis-master:6379",
		Password: "",
		DB:       0,
	})
}

func (r Redis) GetMessage(id uuid.UUID) (message buzzer.Message, err error) {
	messageMap, rError := r.client.HGetAll(id.String()).Result()
	err = rError //explicitly setting error to prevent shadowing

	if err != nil {
		return
	}

	messageId, rError := uuid.Parse(messageMap[buzzer.KeyUuid])
	err = rError //explicitly setting error to prevent shadowing
	messageTime, rError := time.Parse(util.TimeLayout, messageMap[buzzer.KeyBuzzTime])
	err = rError //explicitly setting error to prevent shadowing

	message = buzzer.Message{
		Uuid:       messageId,
		TeamName:   messageMap[buzzer.KeyTeamName],
		PlayerName: messageMap[buzzer.KeyPlayerName],
		BuzzTime:   messageTime,
	}
	return
}

func (r Redis) SetMessage(message buzzer.Message) (messageResult buzzer.Message, err error) {

	message.Uuid, err = uuid.NewRandom()

	if err != nil {
		return

	}

	messageMap := make(map[string]interface{})
	messageMap[buzzer.KeyUuid] = message.Uuid.String()
	messageMap[buzzer.KeyTeamName] = message.TeamName
	messageMap[buzzer.KeyBuzzTime] = message.BuzzTime.Format(util.TimeLayout)
	messageMap[buzzer.KeyPlayerName] = message.PlayerName

	r.client.HMSet(message.Uuid.String(), messageMap)

	return
}

func (r Redis) GetAllMessages() (messages []buzzer.Message, err error) {
	messageKeys, err := r.client.Keys("*").Result()

	for _, messageKey := range messageKeys {
		messageKeyString, rError := uuid.Parse(messageKey)
		err = rError

		if err != nil {
			return
		}

		message, rError := r.GetMessage(messageKeyString)
		err = rError

		if err != nil {
			return
		}

		messages = append(messages, message)
	}

	return
}

func (r Redis) DeleteMessage(id uuid.UUID) (numberOfDeletedKeys int64, err error) {
	numberOfDeletedKeys, err = r.client.Del(id.String()).Result()
	return
}

func (r Redis) DeleteAllMessages() (result string, err error) {
	result, err = r.client.FlushDB().Result()
	return
}

func (r Redis) Ping() (result string, err error) {
	result, err = r.client.Ping().Result()
	return
}
