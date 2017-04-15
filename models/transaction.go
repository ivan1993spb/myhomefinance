package models

import "time"

// TODO create account ID
type Transaction struct {
	ID       uint64
	Time     time.Time
	Amount   float64
	Title    string
	Category string
}
