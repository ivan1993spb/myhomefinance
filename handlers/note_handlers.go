package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/ivan1993spb/myhomefinance/mappers"
)

// NewNoteHandler returns new NoteHandler
func NewNoteHandler(noteMapper mappers.NoteMapper) http.Handler {
	return &noteHandler{noteMapper}
}

type noteHandler struct {
	noteMapper mappers.NoteMapper
}

func (nh *noteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL.Path, 5)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var (
		rawFrom = r.URL.Query().Get("from")
		rawTo   = r.URL.Query().Get("to")
	)
	if len(rawFrom) == 0 || len(rawTo) == 0 {
		log.Println("from or to is empty")
		// log.Println("err", err)
		return
	}
	from, err := time.Parse("2006-Jan-02", rawFrom)
	if err != nil {
		log.Println("err", err)
		return
	}
	to, err := time.Parse("2006-Jan-02", rawTo)
	if err != nil {
		log.Println("err", err)
		return
	}

	log.Println("from", from, "to", to)

	//////////////////////////////////
	notes, err := nh.noteMapper.GetNotesByTimeRange(from, to)
	if err != nil {
		log.Println("err", err)
		return
	}
	/////////////////////////////

	err = json.NewEncoder(w).Encode(notes)
	if err != nil {
		log.Println("err", err)
		return
	}
}
