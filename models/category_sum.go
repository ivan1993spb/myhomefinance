package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type CategorySum struct {
	AccountUUID uuid.UUID
	UserUUID    uuid.UUID
	From        time.Time
	To          time.Time
	Category    string
	Sum         float64
	Count       uint64
}
