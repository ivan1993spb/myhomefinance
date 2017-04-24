package models

import "time"

type CategorySum struct {
	AccountID uint64
	From      time.Time
	To        time.Time
	Category  string
	Sum       float64
	Count     uint64
}
