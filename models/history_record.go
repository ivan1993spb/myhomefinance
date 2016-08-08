package models

import "time"

type HistoryRecord struct {
	DocumentGUID string    `json:"guid"`
	Time         time.Time `json:"time"`
	Name         string    `json:"name"`
	Amount       float64   `json:"amount"`
	Balance      float64   `json:"balance"`
}
