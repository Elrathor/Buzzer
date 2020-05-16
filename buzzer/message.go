package buzzer

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	uuid       uuid.UUID
	TeamName   string
	PlayerName string
	BuzzTime   time.Time
}
