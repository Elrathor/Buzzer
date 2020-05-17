package buzzer

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	Uuid       uuid.UUID
	TeamName   string
	PlayerName string
	BuzzTime   time.Time
}
