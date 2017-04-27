package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/ivan1993spb/myhomefinance/core"
)

type errGetTransactionByIDHandler string

func (e errGetTransactionByIDHandler) Error() string {
	return "error on get transaction by id handler: " + string(e)
}

type GetTransactionByIDHandler struct {
	core *core.Core
	log  *logrus.Logger
}

func NewGetTransactionByIDHandler(core *core.Core, log *logrus.Logger) http.Handler {
	return &GetTransactionByIDHandler{
		core: core,
		log:  log,
	}
}

func (h *GetTransactionByIDHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	transactionID, err := strconv.ParseUint(vars[routeVarTransactionID], 10, 64)
	if err != nil {
		h.log.Error(errGetTransactionByIDHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	transaction, err := h.core.GetTransactionByID(transactionID)
	if err != nil {
		h.log.Error(errGetTransactionByIDHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		h.log.Error(errGetTransactionByIDHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
