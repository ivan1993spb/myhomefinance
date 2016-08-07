package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ivan1993spb/myhomefinance/mappers"
)

// NewGetNotesByTimeRangeHandler returns new getNotesByTimeRangeHandler
func NewGetNotesByTimeRangeHandler(noteMapper mappers.NoteMapper) http.Handler {
	if noteMapper == nil {
		panic(ErrCreateHandlerWithNilMapper)
	}
	return &getNotesByTimeRangeHandler{noteMapper}
}

type getNotesByTimeRangeHandler struct {
	noteMapper mappers.NoteMapper
}

func (h *getNotesByTimeRangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, 5)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	from, to, err := h.getParams(r)
	if err != nil {
		log.Println("err", err)
		return
	}

	notes, err := h.noteMapper.GetNotesByTimeRange(from, to)
	if err != nil {
		log.Println("err", err)
		return
	}

	err = json.NewEncoder(w).Encode(notes)
	if err != nil {
		log.Println("err", err)
		return
	}
}

func (h *getNotesByTimeRangeHandler) getParams(r *http.Request) (time.Time, time.Time, error) {
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
