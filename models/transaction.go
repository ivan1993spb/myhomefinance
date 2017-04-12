package models

import "time"

type Transaction struct {
	ID       uint64
	Time     time.Time
	Amount   float64
	Title    string
	Category string
}
