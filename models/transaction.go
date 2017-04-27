package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type Transaction struct {
	UUID        uuid.UUID
	AccountUUID uuid.UUID
	UserUUID    uuid.UUID
	Time        time.Time
	Amount      float64
	Title       string
	Category    string
}
