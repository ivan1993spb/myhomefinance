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

type errGetUserAccountStatsByTimeRangeHandler string

func (e errGetUserAccountStatsByTimeRangeHandler) Error() string {
	return "error on get user account stats by time range handler: " + string(e)
}

type getUserAccountStatsByTimeRangeHandler struct {
	core *core.Core
	log  *logrus.Logger
}

func NewGetUserAccountStatsByTimeRangeHandler(core *core.Core, log *logrus.Logger) http.Handler {
	return &getUserAccountStatsByTimeRangeHandler{
		core: core,
		log:  log,
	}
}

func (h *getUserAccountStatsByTimeRangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	stats, err := h.core.GetUserAccountStatsByTimeRange(userUUID, accountUUID, from, to)
	if err != nil {
		h.log.Error(errGetUserAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		h.log.Error(errGetUserAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
