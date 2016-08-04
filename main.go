package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tylerb/graceful"

	"fmt"

	"github.com/ivan1993spb/myhomefinance/mappers"
	"github.com/ivan1993spb/myhomefinance/sqlite3mappers"
)

//go:generate go-bindata-assetfs -nometadata -debug -ignore "static/src/" static/...

const (
	URL_PATH_API = "/api"

	URL_PATH_NOTES                         = `/notes`
	URL_PATH_NOTES_ID                      = `/notes/{id}`
	URL_PATH_NOTES_DATE_FROM_DATE_TO       = `/notes/{date_from:\d{4}-\d{2}-\d{2}}_{date_to:\d{4}-\d{2}-\d{2}}`
	URL_PATH_NOTES_DATE_FROM_DATE_TO_MATCH = `/notes/{date_from:\d{4}-\d{2}-\d{2}}_{date_to:\d{4}-\d{2}-\d{2}}/match`

	URL_PATH_HISTORY_DATE_FROM_DATE_TO = `/history/{date_from:\d{4}-\d{2}-\d{2}}_{date_to:\d{4}-\d{2}-\d{2}}`
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

	apiRouter := r.PathPrefix(URL_PATH_API).Subrouter()

	apiRouter.Path(URL_PATH_NOTES_DATE_FROM_DATE_TO_MATCH).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 6)
		noteMapper.GetNotesByTimeRangeMatch(time.Unix(0, 0), time.Now(), "text to match")
	})

	apiRouter.Path(URL_PATH_NOTES_DATE_FROM_DATE_TO).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 5)
		noteMapper.GetNotesByTimeRange(time.Unix(0, 0), time.Now())
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		fmt.Fprint(w, "{\"data\": 123}")
	})

	apiRouter.Path(URL_PATH_NOTES_ID).Methods(http.MethodPut).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 3)
		noteMapper.UpdateNote(1, time.Now(), "new name of note", "new text of note")
	})

	apiRouter.Path(URL_PATH_NOTES_ID).Methods(http.MethodDelete).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 4)
		noteMapper.DeleteNote(1)
	})

	apiRouter.Path(URL_PATH_NOTES).Methods(http.MethodPost).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 1)
		noteMapper.CreateNote(time.Now(), "new note", "any text")
	})

	apiRouter.Path(URL_PATH_NOTES_ID).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
