package handlers

import (
	"fmt"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"

	"github.com/ivan1993spb/myhomefinance/core"
)

const URLRouteDeleteTransaction = "/user/{" + routeVarUserUUID + "}/account/{" + routeVarAccountUUID + "}/transaction/{" + routeVarTransactionUUID + "}"

const formatURLRouteDeleteTransaction = "/user/%s/account/%s/transaction/%s"

func BuildPathDeleteTransaction(userUUID, accountUUID, transactionUUID uuid.UUID) string {
	return fmt.Sprintf(formatURLRouteDeleteTransaction, userUUID, accountUUID, transactionUUID)
}

type errDeleteTransactionHandler string

func (e errDeleteTransactionHandler) Error() string {
	return "error on delete transaction handler: " + string(e)
}

type deleteTransactionHandler struct {
	core *core.Core
	log  *logrus.Logger
}

func NewDeleteTransactionHandler(core *core.Core, log *logrus.Logger) http.Handler {
	return &deleteTransactionHandler{
		core: core,
		log:  log,
	}
}

func (h *deleteTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userUUID, err := uuid.FromString(vars[routeVarUserUUID])
	if err != nil {
		h.log.Error(errCreateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	accountUUID, err := uuid.FromString(vars[routeVarAccountUUID])
	if err != nil {
		h.log.Error(errCreateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	transactionUUID, err := uuid.FromString(vars[routeVarTransactionUUID])
	if err != nil {
		h.log.Error(errDeleteTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = h.core.DeleteTransaction(userUUID, accountUUID, transactionUUID)
	if err != nil {
		h.log.Error(errDeleteTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
