package sqlite3mappers

import (
	"database/sql"
	"time"

	"github.com/ivan1993spb/myhomefinance/mappers"
	"github.com/ivan1993spb/myhomefinance/models"
)

const _SQL_SELECT_HISTORY_RECORDS_BY_TIME_RANGE = `
SELECT t1.document_guid, t1.unixtimestamp, t1.name, t1.amount, ROUND(SUM(t2.amount), 2) AS balance
FROM (
    SELECT document_guid, unixtimestamp, name, amount, description
    FROM inflow
    WHERE unixtimestamp BETWEEN $1 AND $2

    UNION

    SELECT document_guid, unixtimestamp, name, -amount AS amount, description FROM outflow
        WHERE unixtimestamp BETWEEN $1 AND $2
) AS t1,
(
    SELECT document_guid, unixtimestamp, name, amount, description FROM inflow
    WHERE unixtimestamp BETWEEN $1 AND $2

    UNION

    SELECT document_guid, unixtimestamp, name, -amount AS amount, description FROM outflow
    WHERE unixtimestamp BETWEEN $1 AND $2
) AS t2
WHERE t2.unixtimestamp <= t1.unixtimestamp
GROUP BY t1.document_guid
ORDER BY t1.unixtimestamp DESC
`

type HistoryRecordMapper struct {
	db *sql.DB

	selectHistoryRecordsByTimeRange *sql.Stmt
}

func NewHistoryRecordMapper(db *sql.DB) (*HistoryRecordMapper, error) {
	err := db.Ping()
	if err != nil {
		return nil, mappers.ErrCreateHistoryRecordMapper(err.Error())
	}

	historyRecordMapper := &HistoryRecordMapper{db: db}

	historyRecordMapper.selectHistoryRecordsByTimeRange, err = db.Prepare(_SQL_SELECT_HISTORY_RECORDS_BY_TIME_RANGE)
	if err != nil {
		return nil, mappers.ErrCreateHistoryRecordMapper(err.Error())
	}

	return historyRecordMapper, err
}

func (hrm *HistoryRecordMapper) GetHistoryRecordsByTimeRange(from time.Time, to time.Time) (
	[]*models.HistoryRecord, error) {

	if from.Unix() >= to.Unix() {
		return nil, mappers.ErrGetHistoryRecordsByTimeRangeInvalidTimeRange
	}

	rows, err := hrm.selectHistoryRecordsByTimeRange.Query(from.Unix(), to.Unix())
	if err != nil {
		return nil, mappers.ErrGetHistoryRecordsByTimeRange("cannot select history records: " + err.Error())
	}
	defer rows.Close()

	historyRecords := make([]*models.HistoryRecord, 0)

	for rows.Next() {
		var unixtimestamp int64
		hr := &models.HistoryRecord{}
		if err := rows.Scan(&hr.DocumentGUID, &unixtimestamp, &hr.Name, &hr.Amount, &hr.Balance); err != nil {
			return nil, mappers.ErrGetHistoryRecordsByTimeRange("error on scanning query result: " + err.Error())
		}
		hr.Time = time.Unix(unixtimestamp, 0)
		historyRecords = append(historyRecords, hr)
	}

	if err := rows.Err(); err != nil {
		return nil, mappers.ErrGetHistoryRecordsByTimeRange("query error: " + err.Error())
	}

	return historyRecords, nil
}
