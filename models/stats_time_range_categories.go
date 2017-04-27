package models

import (
	"time"

	"github.com/satori/go.uuid"
)

type StatsTimeRangeCategories struct {
	AccountUUID uuid.UUID
	UserUUID    uuid.UUID
	From        time.Time
	To          time.Time
	Inflow      float64
	Outflow     float64
	Profit      float64
	Count       uint64
	Categories  []string
}
