package handlers

import (
	"net/http"
	"time"

	"github.com/ivan1993spb/myhomefinance/mappers"
)

// NewCreateNoteHandler returns new createNoteHandler handler
func NewCreateNoteHandler(noteMapper mappers.NoteMapper) http.Handler {
	if noteMapper == nil {
		panic(ErrCreateHandlerWithNilMapper)
	}
	return &createNoteHandler{noteMapper}
}

type createNoteHandler struct {
	noteMapper mappers.NoteMapper
}

func (h *createNoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// h.noteMapper.CreateNote(time, name, text)
}

func (h *createNoteHandler) getParams(r *http.Request) (time.Time, string, string) {
	return time.Time{}, "", ""
}
