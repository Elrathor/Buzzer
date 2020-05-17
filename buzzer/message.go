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
	Uuid       uuid.UUID
	TeamName   string
	PlayerName string
	BuzzTime   time.Time
}
