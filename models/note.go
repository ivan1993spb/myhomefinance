package models

import "time"

type Note struct {
	Id   uint64
	Time time.Time
	Name string
	Text string
}
