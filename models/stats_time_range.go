package models

import "time"

type StatsTimeRange struct {
	AccountID uint64
	From      time.Time
	To        time.Time
	Inflow    float64
	Outflow   float64
	Profit    float64
	Count     uint64
}
