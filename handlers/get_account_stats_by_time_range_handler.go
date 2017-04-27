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

type errGetAccountStatsByTimeRangeHandler string

func (e errGetAccountStatsByTimeRangeHandler) Error() string {
	return "error on get account stats by time range handler: " + string(e)
}

type getAccountStatsByTimeRangeHandler struct {
	core *core.Core
	log  *logrus.Logger
}

func NewGetAccountStatsByTimeRangeHandler(core *core.Core, log *logrus.Logger) http.Handler {
	return &getAccountStatsByTimeRangeHandler{
		core: core,
		log:  log,
	}
}

func (h *getAccountStatsByTimeRangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	stats, err := h.core.GetUserAccountStatsByTimeRange(accountID, from, to)
	if err != nil {
		h.log.Error(errGetAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(stats)
	if err != nil {
		h.log.Error(errGetAccountStatsByTimeRangeHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
