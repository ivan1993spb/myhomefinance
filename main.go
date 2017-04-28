package main

import (
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/tylerb/graceful"
	"github.com/urfave/negroni"

	"github.com/ivan1993spb/myhomefinance/core"
	"github.com/ivan1993spb/myhomefinance/handlers"
	"github.com/ivan1993spb/myhomefinance/memoryrepository"
	"github.com/ivan1993spb/myhomefinance/middlewares"
)

func main() {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()

	r := mux.NewRouter()

	r.Path(handlers.URLRouteCreateTransaction).Methods(http.MethodPost).Handler(handlers.NewCreateTransactionHandler(c, log))
	r.Path(handlers.URLRouteUpdateTransaction).Methods(http.MethodPut).Handler(handlers.NewUpdateTransactionHandler(c, log))
	r.Path(handlers.URLRouteDeleteTransaction).Methods(http.MethodDelete).Handler(handlers.NewDeleteTransactionHandler(c, log))

	n := negroni.Classic()
	n.Use(middlewares.Secure)
	n.UseHandler(r)

	(&graceful.Server{
		Server: &http.Server{Addr: ":8888", Handler: n},
		BeforeShutdown: func() bool {
			return true
		},
	}).ListenAndServe()
}
