package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tylerb/graceful"

	"github.com/ivan1993spb/myhomefinance/mappers"
	"github.com/ivan1993spb/myhomefinance/sqlite3mappers"
)

//go:generate go-bindata-assetfs -nometadata -debug -ignore "static/src/" static/...

const (
	urlPathAPI = "/api"

	urlPathNote            = `/note`
	urlPathNotesRange      = `/notes/range`
	urlPathNotesRangeMatch = `/notes/range/match`
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

	apiRouter := r.PathPrefix(urlPathAPI).Subrouter()

	apiRouter.Path(urlPathNotesRangeMatch).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 6)
		noteMapper.GetNotesByTimeRangeMatch(time.Unix(0, 0), time.Now(), "text to match")
	})

	apiRouter.Path(urlPathNotesRange).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 5)
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var (
			rawFrom = r.URL.Query().Get("from")
			rawTo   = r.URL.Query().Get("to")
		)
		if len(rawFrom) == 0 || len(rawTo) == 0 {
			log.Println("from or to is empty")
			log.Println("err", err)
			return
		}
		from, err := time.Parse("2016-Jan-22", rawFrom)
		if err != nil {
			log.Println("err", err)
			return
		}
		to, err := time.Parse("2016-Jan-22", rawTo)
		if err != nil {
			log.Println("err", err)
			return
		}

		//////////////////////////////////
		notes, err := noteMapper.GetNotesByTimeRange(from, to)
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
	})

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

	r.PathPrefix("/").Methods(http.MethodGet).Handler(http.FileServer(assetFS()))

	(&graceful.Server{
		Server: &http.Server{Addr: ":8888", Handler: r},
		BeforeShutdown: func() bool {
			db.Close()
			return true
		},
	}).ListenAndServe()

}
