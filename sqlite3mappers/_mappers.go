package sqlite3mappers

/*
import (
	"database/sql"

	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/ivan1993spb/myhomefinance/models"
	_ "github.com/mattn/go-sqlite3"
)

type ErrInflowMapper string

func (e ErrInflowMapper) Error() string {
	return "inflow mapper error: " + string(e)
}

type InflowMapper struct {
	*sql.DB
}

// CreateInflow creates new inflow document into db and returns it
func (im *InflowMapper) CreateInflow(t time.Time, name string, amount float64, description, source string,
) (*models.Inflow, error) {
	if len(name) == 0 {
		return nil, ErrInflowMapper("name cannot be empty")
	}
	if amount <= 0 {
		return nil, ErrInflowMapper(fmt.Sprintf("invalid amount %d (must be > 0)", amount))
	}

	guid, err := newGUID()
	if err != nil {
		return nil, ErrInflowMapper("cannot generate guid: " + err.Error())
	}

	res, err := im.DB.Exec("INSERT INTO `inflow` (`document_guid`, `unixtimestamp`, `name`, `amount`, "+
		"`description`, `source`) VALUES(?, ?, ?, ?, ?, ?)", guid, t.Unix(), name, amount, description,
		source)
	if err != nil {
		return nil, ErrInflowMapper("cannot insert new inflow into db: " + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, ErrInflowMapper("cannot get id of new inflow: " + err.Error())
	}

	return &models.Inflow{
		ID:          id,
		Datetime:    strfmt.DateTime(t),
		Name:        &name,
		Amount:      &amount,
		Description: description,
		Source:      &source,
	}, nil
}

type ErrOutflowMapper string

func (e ErrOutflowMapper) Error() string {
	return "outflow mapper error: " + string(e)
}

type OutflowMapper struct {
	*sql.DB
}

// CreateOutflow creates new outflow document into db and returns it
func (om *OutflowMapper) CreateOutflow(t time.Time, name string, amount float64, description, destination,
	target string, count float64, metricUnit string, satisfaction float32) (*models.Outflow, error) {
	if len(name) == 0 {
		return nil, ErrOutflowMapper("name cannot be empty")
	}
	if amount <= 0 {
		return nil, ErrOutflowMapper(fmt.Sprintf("invalid amount %d (must be > 0)", amount))
	}

	guid, err := newGUID()
	if err != nil {
		return nil, ErrOutflowMapper("cannot generate guid: " + err.Error())
	}

	res, err := om.DB.Exec("INSERT INTO `outflow` (`document_guid`, `unixtimestamp`, `name`, `amount`, "+
		"`description`, `destination`, `target`, `count`, `metric_unit`, `satisfaction`) VALUES "+
		"(?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		guid, t.Unix(), name, amount, description, description, target, count, metricUnit, satisfaction)
	if err != nil {
		return nil, ErrOutflowMapper("cannot insert new outflow into db: " + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, ErrOutflowMapper("cannot get id of new outflow: " + err.Error())
	}

	return &models.Outflow{
		ID:           id,
		Datetime:     &strfmt.DateTime(t),
		Name:         &name,
		Amount:       &amount,
		Description:  description,
		Destination:  &destination,
		Target:       target,
		Count:        count,
		MetricUnit:   metricUnit,
		Satisfaction: satisfaction,
	}, nil
}



*/
