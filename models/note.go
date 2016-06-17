package models

import "time"

type Note struct {
	Id   int64
	Time time.Time
	Name string
	Text string
}
