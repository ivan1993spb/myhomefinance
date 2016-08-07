package models

import "time"

type Note struct {
	Id   int64     `json:"id"`
	Time time.Time `json:"time"`
	Name string    `json:"name"`
	Text string    `json:"text,omitempty"`
}
