package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"

	"github.com/ivan1993spb/myhomefinance/core"
)

type errGetUserAccountTransactionsByTimeRange string

func (e errGetUserAccountTransactionsByTimeRange) Error() string {
	return "error on get user account transactions by time range handler: " + string(e)
}

type getUserAccountTransactionsByTimeRangeHandler struct {
	core *core.Core
	log  *logrus.Logger
}

func NewGetUserAccountTransactionsByTimeRangeHandler(core *core.Core, log *logrus.Logger) http.Handler {
	return &getUserAccountTransactionsByTimeRangeHandler{
		core: core,
		log:  log,
	}
}

func (h *getUserAccountTransactionsByTimeRangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
	from, err := time.Parse(apiDateFormat, r.URL.Query().Get(fieldTimeFrom))
	if err != nil {
		h.log.Error(errGetUserAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	to, err := time.Parse(apiDateFormat, r.URL.Query().Get(fieldTimeTo))
	if err != nil {
		h.log.Error(errGetUserAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	transactions, err := h.core.GetUserAccountTransactionsByTimeRange(userUUID, accountUUID, from, to)
	if err != nil {
		h.log.Error(errGetUserAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(transactions)
	if err != nil {
		h.log.Error(errGetUserAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
