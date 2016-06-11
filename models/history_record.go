package main

import "time"

type HistoryRecord struct {
	DocumentGUID string
	Time         time.Time
	Name         string
	Amount       float64
	Balance      float64
}
