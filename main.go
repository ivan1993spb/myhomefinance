package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tylerb/graceful"

	"github.com/ivan1993spb/myhomefinance/mappers"
	"github.com/ivan1993spb/myhomefinance/sqlite3mappers"
)

const (
	URL_PATH_NOTES                   = `/notes`
	URL_PATH_NOTES_ID                = `/notes/{id}`
	URL_PATH_DATE_FROM_DATE_TO       = `/notes/{date_from:\d{4}-\d{2}-\d{2}}_{date_to:\d{4}-\d{2}-\d{2}}`
	URL_PATH_DATE_FROM_DATE_TO_MATCH = `/notes/{date_from:\d{4}-\d{2}-\d{2}}_{date_to:\d{4}-\d{2}-\d{2}}/match`
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

	r.Path(URL_PATH_NOTES).Methods(http.MethodPost).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 1)
		noteMapper.CreateNote(time.Now(), "new note", "any text")
	})
	r.Path(URL_PATH_NOTES_ID).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 2)
		noteMapper.GetNoteById(1)
	})
	r.Path(URL_PATH_NOTES_ID).Methods(http.MethodPut).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 3)
		noteMapper.UpdateNote(1, time.Now(), "new name of note", "new text of note")
	})
	r.Path(URL_PATH_NOTES_ID).Methods(http.MethodDelete).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 4)
		noteMapper.DeleteNote(1)
	})
	r.Path(URL_PATH_DATE_FROM_DATE_TO).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 5)
		noteMapper.GetNotesByTimeRange(time.Unix(0, 0), time.Now())
	})
	r.Path(URL_PATH_DATE_FROM_DATE_TO_MATCH).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 6)
		noteMapper.GetNotesByTimeRangeMatch(time.Unix(0, 0), time.Now(), "text to match")
	})

	(&graceful.Server{
		Server: &http.Server{Addr: ":8888", Handler: r},
		BeforeShutdown: func() {
			db.Close()
		},
	}).ListenAndServe()
}
