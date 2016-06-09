package sqlite3mappers

import (
	"database/sql"
	"os"

	"fmt"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/ivan1993spb/myhomefinance/models"
	_ "github.com/mattn/go-sqlite3"
)

//go:generate go-bindata --pkg=sqlite3mappers db.sql

type ErrSQLiteDB string

func (e ErrSQLiteDB) Error() string {
	return "sqlite db error: " + string(e)
}

// InitSQLiteDB tries to load sqlite db from file or creates new db file with tables and views
func InitSQLiteDB(dbFileName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFileName)
	if err != nil {
		return db, ErrSQLiteDB(err.Error())
	}

	if _, err := os.Stat(dbFileName); os.IsNotExist(err) {
		_, err = db.Exec(string(MustAsset("db.sql")))
		if err != nil {
			return db, ErrSQLiteDB("cannot execute queries for tables and views creation: " +
				err.Error())
		}
	}

	return db, err
}

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

type ErrHistoryMapper string

func (e ErrHistoryMapper) Error() string {
	return "history mapper error:" + string(e)
}

type HistoryMapper struct {
	*sql.DB
}

func (hm *HistoryMapper) GetHistoryFeed(from, to time.Time) ([]*HistoryRecord, error) {
	if from.Unix() >= to.Unix() {
		return nil, ErrHistoryMapper("invalid time range")
	}

	rows, err := hm.DB.Query("SELECT `t1`.`document_guid`, `t1`.`unixtimestamp`, `t1`.`name`, `t1`.`amount`,"+
		"    SUM(`t2`.`amount`) AS `balance`"+
		"    FROM ("+
		"        SELECT `document_guid`, `unixtimestamp`, `name`, `amount`, `description` FROM `inflow`"+
		"            WHERE `unixtimestamp` BETWEEN $1 AND $2"+
		"        UNION"+
		"        SELECT `document_guid`, `unixtimestamp`, `name`, -`amount` AS `amount`, `description` FROM `outflow`"+
		"            WHERE `unixtimestamp` BETWEEN $1 AND $2"+
		"    ) AS `t1`,"+
		"    ("+
		"        SELECT `document_guid`, `unixtimestamp`, `name`, `amount`, `description` FROM `inflow`"+
		"            WHERE `unixtimestamp` BETWEEN $1 AND $2"+
		"        UNION"+
		"        SELECT `document_guid`, `unixtimestamp`, `name`, -`amount` AS `amount`, `description` FROM `outflow`"+
		"            WHERE `unixtimestamp` BETWEEN $1 AND $2"+
		"    ) AS `t2`"+
		"        WHERE `t2`.`unixtimestamp` <= `t1`.`unixtimestamp`"+
		"    GROUP BY `t1`.`document_guid` ORDER BY `t1`.`unixtimestamp` DESC", from.Unix(), to.Unix())
	if err != nil {
		return nil, ErrHistoryMapper("cannot select history records: " + err.Error())
	}
	defer rows.Close()

	historyRecords := make([]*HistoryRecord, 0)

	for rows.Next() {
		var unixtimestamp int64
		hr := &HistoryRecord{}
		if err := rows.Scan(&hr.DocumentGUID, &unixtimestamp, &hr.Name, &hr.Amount, &hr.Balance); err != nil {
			return nil, ErrHistoryMapper("error on scanning query result: " + err.Error())
		}
		hr.Time = time.Unix(unixtimestamp, 0)
		historyRecords = append(historyRecords, hr)
	}

	if err := rows.Err(); err != nil {
		return nil, ErrHistoryMapper("query error: " + err.Error())
	}

	return historyRecords, nil
}
