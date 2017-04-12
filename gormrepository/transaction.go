package gormrepository

import "time"

type transaction struct {
	ID       uint64
	Time     time.Time
	Amount   float64
	Title    string
	Category string
}
