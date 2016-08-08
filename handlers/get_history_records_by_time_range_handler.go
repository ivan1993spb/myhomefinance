package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ivan1993spb/myhomefinance/mappers"
)

// NewGetHistoryRecordsByTimeRangeHandler returns new getHistoryRecordsByTimeRangeHandler
func NewGetHistoryRecordsByTimeRangeHandler(historyRecordMapper mappers.HistoryRecordMapper) http.Handler {
	if historyRecordMapper == nil {
		panic(ErrCreateHandlerWithNilMapper)
	}
	return &getHistoryRecordsByTimeRangeHandler{historyRecordMapper}
}

type getHistoryRecordsByTimeRangeHandler struct {
	historyRecordMapper mappers.HistoryRecordMapper
}

func (h *getHistoryRecordsByTimeRangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	from, to, err := h.getParams(r)
	if err != nil {
		log.Println("err", err)
		return
	}

	historyRecords, err := h.historyRecordMapper.GetHistoryRecordsByTimeRange(from, to)
	if err != nil {
		log.Println("err", err)
		return
	}

	err = json.NewEncoder(w).Encode(historyRecords)
	if err != nil {
		log.Println("err", err)
		return
	}
}

func (h *getHistoryRecordsByTimeRangeHandler) getParams(r *http.Request) (time.Time, time.Time, error) {
	rawFrom := r.URL.Query().Get("from")
	if len(rawFrom) == 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("received empty date from")
	}

	rawTo := r.URL.Query().Get("to")
	if len(rawTo) == 0 {
		return time.Time{}, time.Time{}, fmt.Errorf("received empty date to")
	}

	from, err := time.Parse(apiDateFormat, rawFrom)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("cannot parse date from: %s", err)
	}

	to, err := time.Parse(apiDateFormat, rawTo)
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("cannot parse date to: %s", err)
	}

	return from, to, nil
}
