package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"

	"github.com/ivan1993spb/myhomefinance/core"
)

type errGetUserAccountTransactionHandler string

func (e errGetUserAccountTransactionHandler) Error() string {
	return "error on get user account transaction handler: " + string(e)
}

type getUserAccountTransactionHandler struct {
	core *core.Core
	log  *logrus.Logger
}

func NewGetTransactionByIDHandler(core *core.Core, log *logrus.Logger) http.Handler {
	return &getUserAccountTransactionHandler{
		core: core,
		log:  log,
	}
}

func (h *getUserAccountTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		h.log.Error(errUpdateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	transaction, err := h.core.GetUserAccountTransaction(userUUID, accountUUID, transactionUUID)
	if err != nil {
		h.log.Error(errGetUserAccountTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		h.log.Error(errGetUserAccountTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
