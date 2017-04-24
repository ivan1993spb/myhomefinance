package models

import "time"

type StatsTimeRangeCategories struct {
	AccountID  uint64
	From       time.Time
	To         time.Time
	Inflow     float64
	Outflow    float64
	Profit     float64
	Count      uint64
	Categories []string
}
