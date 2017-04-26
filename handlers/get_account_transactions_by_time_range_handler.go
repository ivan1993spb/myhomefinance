package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/ivan1993spb/myhomefinance/core"
)

type errGetAccountTransactionsByTimeRangeHandler string

func (e errGetAccountTransactionsByTimeRangeHandler) Error() string {
	return "error on get account transactions by time range handler: " + string(e)
}

type getAccountTransactionsByTimeRangeHandler struct {
	core *core.Core
	log  *logrus.Logger
}

func NewGetAccountTransactionsByTimeRangeHandler (core *core.Core, log *logrus.Logger) http.Handler {
	return &getAccountTransactionsByTimeRangeHandler{
		core: core,
		log:  log,
	}
}

func (h *getAccountTransactionsByTimeRangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	accountID, err := strconv.ParseUint(vars[routeVarAccountID], 10, 64)
	if err != nil {
		h.log.Error(errGetAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	from, err := time.Parse(apiDateFormat, r.URL.Query().Get(fieldTimeFrom))
	if err != nil {
		h.log.Error(errGetAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	to, err := time.Parse(apiDateFormat, r.URL.Query().Get(fieldTimeTo))
	if err != nil {
		h.log.Error(errGetAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	transactions, err := h.core.GetAccountTransactionsByTimeRange(accountID, from, to)
	if err != nil {
		h.log.Error(errGetAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		h.log.Error(errGetAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
