package models

import "time"

type Outflow struct {
	Id           int64
	DocumentGUID string
	Time         time.Time
	Name         string
	Amount       float64
	Description  string
	Destination  string
	Target       string
	Count        float64
	MetricUnit   string
	Satisfaction float32
}
