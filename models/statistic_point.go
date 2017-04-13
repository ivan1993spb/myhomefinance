package models

import "time"

type StatisticPoint struct {
	ID   uint64
	Time time.Time

	Inflow           float64
	Outflow          float64
	Profit           float64
	Balance          float64
	MeanDaylyOutflow float64
}
