package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/tylerb/graceful"
)

const (
	URL_PATH_NOTES                   = `/notes`
	URL_PATH_NOTES_ID                = `/notes/{id}`
	URL_PATH_DATE_FROM_DATE_TO       = `/notes/{date_from:\d{4}-\d{2}-\d{2}}_{date_to:\d{4}-\d{2}-\d{2}}`
	URL_PATH_DATE_FROM_DATE_TO_MATCH = `/notes/{date_from:\d{4}-\d{2}-\d{2}}_{date_to:\d{4}-\d{2}-\d{2}}/match`
)

func main() {
	r := mux.NewRouter()

	r.Path(URL_PATH_NOTES).Methods(http.MethodPost).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 3)
	})
	r.Path(URL_PATH_NOTES_ID).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 1)
	})
	r.Path(URL_PATH_NOTES_ID).Methods(http.MethodPut).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 2)
	})
	r.Path(URL_PATH_NOTES_ID).Methods(http.MethodDelete).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 4)
	})
	r.Path(URL_PATH_DATE_FROM_DATE_TO).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 5)
	})
	r.Path(URL_PATH_DATE_FROM_DATE_TO_MATCH).Methods(http.MethodGet).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, 6)
	})

	graceful.Run(":8888", time.Second, r)
}
