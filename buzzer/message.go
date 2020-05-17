package buzzer

import (
	"github.com/google/uuid"
	"time"
)

const (
	KeyTeamName   = "TeamName"
	KeyPlayerName = "PlayerName"
	KeyBuzzTime   = "BuzzTime"
	KeyUuid       = "Uuid"
)

type Message struct {
	Uuid       uuid.UUID `form:"Uuid" json:"Uuid" xml:"Uuid"`
	TeamName   string `form:"TeamName" json:"TeamName" xml:"TeamName"  binding:"required"`
	PlayerName string `form:"PlayerName" json:"PlayerName" xml:"PlayerName"  binding:"required"`
	BuzzTime   time.Time `form:"BuzzTime" json:"BuzzTime" xml:"BuzzTime"  binding:"required"`
}
