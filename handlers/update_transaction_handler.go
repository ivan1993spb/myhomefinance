package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"

	"github.com/ivan1993spb/myhomefinance/core"
)

type errUpdateTransactionHandler string

func (e errUpdateTransactionHandler) Error() string {
	return "error on update transaction handler: " + string(e)
}

type updateTransactionHandler struct {
	core *core.Core
	log  *logrus.Logger
}

func NewUpdateTransactionHandler(core *core.Core, log *logrus.Logger) http.Handler {
	return &updateTransactionHandler{
		core: core,
		log:  log,
	}
}

func (h *updateTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	transactionID, err := strconv.ParseUint(vars[routeVarTransactionID], 10, 64)
	if err != nil {
		h.log.Error(errUpdateTransactionHandler(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	amount, err := strconv.ParseFloat(r.PostFormValue(fieldAmount), 64)
	if err != nil {
		h.log.Error(errUpdateTransactionHandler(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	t, err := time.Parse(apiDateFormat, r.PostFormValue(fieldTime))
	if err != nil {
		h.log.Error(errUpdateTransactionHandler(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	title := r.PostFormValue(fieldTitle)
	category := r.PostFormValue(fieldCategory)

	err = h.core.UpdateTransaction(transactionID, t, amount, title, category)
	if err != nil {
		h.log.Error(errUpdateTransactionHandler(err))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
}
