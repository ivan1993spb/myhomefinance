package main

import "time"

type Inflow struct {
	Id           uint64
	DocumentGUID string
	Time         time.Time
	Name         string
	Amount       float64
	Description  string
	Source       string
}
