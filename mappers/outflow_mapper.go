package mappers

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

type OutflowMapper interface {
	CreateOutflow(t time.Time, name string, amount float64, description, destination, target string, count float64,
		metricUnit string, satisfaction float32) (*models.Outflow, error)
}
