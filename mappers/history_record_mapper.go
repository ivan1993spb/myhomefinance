package mappers

import (
	"time"

	"github.com/ivan1993spb/myhomefinance/models"
)

type HistoryRecordMapper interface {
	GetHistoryRecordsByTimeRange(from time.Time, to time.Time) ([]*models.HistoryRecord, error)
}

type ErrCreateHistoryRecordMapper string

func (e ErrCreateHistoryRecordMapper) Error() string {
	return "cannot create history record mapper: " + string(e)
}

type ErrGetHistoryRecordsByTimeRange string

func (e ErrGetHistoryRecordsByTimeRange) Error() string {
	return "cannot get history records by time range: " + string(e)
}

var ErrGetHistoryRecordsByTimeRangeInvalidTimeRange = ErrGetHistoryRecordsByTimeRange("invalid time range")
