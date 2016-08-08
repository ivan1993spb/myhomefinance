package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tylerb/graceful"

	"github.com/ivan1993spb/myhomefinance/handlers"
	"github.com/ivan1993spb/myhomefinance/mappers"
	"github.com/ivan1993spb/myhomefinance/sqlite3mappers"
)

//go:generate go-bindata-assetfs -nometadata -debug -ignore "static/src/" static/...

const (
	urlPathAPI = "/api"

	urlPathNote            = `/note`
	urlPathNotesRange      = `/notes/range`
	urlPathNotesRangeMatch = `/notes/range/match`

	urlPathHistoryRange = `/history/range`
)

func initDb() (*sql.DB, error) {
	return sqlite3mappers.InitSQLiteDB("test.db")
}

func main() {
	r := mux.NewRouter()
	db, err := initDb()
	if err != nil {
		log.Println(err)
	}
	var noteMapper mappers.NoteMapper
	noteMapper, err = sqlite3mappers.NewNoteMapper(db)

	var historyRecordMapper mappers.HistoryRecordMapper
	historyRecordMapper, err = sqlite3mappers.NewHistoryRecordMapper(db)
	apiRouter := r.PathPrefix(urlPathAPI).Subrouter()

	apiRouter.Path(urlPathNotesRangeMatch).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 6)
		noteMapper.GetNotesByTimeRangeMatch(time.Unix(0, 0), time.Now(), "text to match")
	})

	apiRouter.Path(urlPathNotesRange).Methods(http.MethodGet).Handler(handlers.NewGetNotesByTimeRangeHandler(noteMapper))

	apiRouter.Path(urlPathNote).Methods(http.MethodPut).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 3)
		noteMapper.UpdateNote(1, time.Now(), "new name of note", "new text of note")
	})

	apiRouter.Path(urlPathNote).Methods(http.MethodDelete).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 4)
		noteMapper.DeleteNote(1)
	})

	apiRouter.Path(urlPathNote).Methods(http.MethodPost).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 1)
		noteMapper.CreateNote(time.Now(), "new note", "any text")
	})

	apiRouter.Path(urlPathNote).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 2)
		noteMapper.GetNoteById(1)
	})

	apiRouter.Path(urlPathHistoryRange).Methods(http.MethodGet).Handler(handlers.NewGetHistoryRecordsByTimeRangeHandler(historyRecordMapper))

	r.PathPrefix("/").Methods(http.MethodGet).Handler(http.FileServer(assetFS()))

	(&graceful.Server{
		Server: &http.Server{Addr: ":8888", Handler: r},
		BeforeShutdown: func() bool {
			db.Close()
			return true
		},
	}).ListenAndServe()

}
