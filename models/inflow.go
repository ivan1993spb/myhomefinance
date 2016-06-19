package models

import "time"

type Inflow struct {
	Id           int64
	DocumentGUID string
	Time         time.Time
	Name         string
	Amount       float64
	Description  string
	Source       string
}
