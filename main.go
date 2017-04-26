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
)

func main() {
	userRepository, _ := memoryrepository.NewUserRepository()
	accountRepository, _ := memoryrepository.NewAccountRepository()
	transactionRepository, _ := memoryrepository.NewTransactionRepository()
	c := core.New(userRepository, accountRepository, transactionRepository)
	log := logrus.New()

	r := mux.NewRouter()
	apiRouter := r.PathPrefix("/").Subrouter()

	apiRouter.Path("/account/{account_id}/transaction").Methods(http.MethodPost).Handler(handlers.NewCreateTransactionHandler(c, log))
	apiRouter.Path("/account/{account_id}/transaction/{transaction_id}").Methods(http.MethodPut).Handler(handlers.NewUpdateTransactionHandler(c, log))
	apiRouter.Path("/account/{account_id}/transaction/{transaction_id}").Methods(http.MethodDelete).Handler(handlers.NewDeleteTransactionHandler(c, log))

	n := negroni.Classic()
	n.UseHandler(apiRouter)

	(&graceful.Server{
		Server: &http.Server{Addr: ":8888", Handler: n},
		BeforeShutdown: func() bool {
			return true
		},
	}).ListenAndServe()
}
