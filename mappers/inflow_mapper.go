package mappers

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

type InflowMapper interface {
	CreateInflow(t time.Time, name string, amount float64, description, source string) (*models.Inflow, error)
}
