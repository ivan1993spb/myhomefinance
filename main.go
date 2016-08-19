package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"

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
	db, err := initDb()
	if err != nil {
		log.Println(err)
	}

	var noteMapper mappers.NoteMapper
	noteMapper, err = sqlite3mappers.NewNoteMapper(db)
	if err != nil {
		panic(err)
	}

	var historyRecordMapper mappers.HistoryRecordMapper
	historyRecordMapper, err = sqlite3mappers.NewHistoryRecordMapper(db)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	apiRouter := r.PathPrefix(urlPathAPI).Subrouter()

	apiRouter.Path(urlPathNotesRangeMatch).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noteMapper.GetNotesByTimeRangeMatch(time.Unix(0, 0), time.Now(), "text to match")
	})

	apiRouter.Path(urlPathNotesRange).Methods(http.MethodGet).Handler(handlers.NewGetNotesByTimeRangeHandler(noteMapper))

	apiRouter.Path(urlPathNote).Methods(http.MethodPut).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noteMapper.UpdateNote(1, time.Now(), "new name of note", "new text of note")
	})

	apiRouter.Path(urlPathNote).Methods(http.MethodDelete).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noteMapper.DeleteNote(1)
	})

	apiRouter.Path(urlPathNote).Methods(http.MethodPost).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noteMapper.CreateNote(time.Now(), "new note", "any text")
	})

	apiRouter.Path(urlPathNote).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noteMapper.GetNoteById(1)
	})

	apiRouter.Path(urlPathHistoryRange).Methods(http.MethodGet).Handler(handlers.NewGetHistoryRecordsByTimeRangeHandler(historyRecordMapper))

	apiJSON := negroni.New()
	apiJSON.UseFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(rw, r)
	})
	apiJSON.UseHandler(r)

	static := negroni.NewStatic(assetFS())

	n := negroni.Classic()
	n.UseFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		switch r.URL.Path {
		case "/notes", "/history", "/graphs":
			r.URL.Path = "/index.html"
		}

		next(rw, r)
	})
	n.Use(static)
	n.UseHandler(apiJSON)

	(&graceful.Server{
		Server: &http.Server{Addr: ":8888", Handler: n},
		BeforeShutdown: func() bool {
			db.Close()
			return true
		},
	}).ListenAndServe()
}
