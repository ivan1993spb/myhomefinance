package handlers

import (
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/ivan1993spb/myhomefinance/core"
)

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

	transactionID, err := strconv.ParseUint(vars[routeVarTransactionID], 10, 64)
	if err != nil {
		h.log.Error(errDeleteTransactionHandler(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = h.core.DeleteTransaction(transactionID)
	if err != nil {
		h.log.Error(errDeleteTransactionHandler(err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
