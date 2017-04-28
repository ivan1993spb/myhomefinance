package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"

	"github.com/ivan1993spb/myhomefinance/core"
)

const URLRouteUpdateTransaction = "/user/{" + routeVarUserUUID + "}/account/{" + routeVarAccountUUID + "}/transaction/{" + routeVarTransactionUUID + "}"

const formatURLRouteUpdateTransaction = "/user/%s/account/%s/transaction/%s"

func BuildPathUpdateTransaction(userUUID, accountUUID, transactionUUID uuid.UUID) string {
	return fmt.Sprintf(formatURLRouteUpdateTransaction, userUUID, accountUUID, transactionUUID)
}

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

	amount, err := strconv.ParseFloat(r.PostFormValue(fieldAmount), 64)
	if err != nil {
		h.log.Error(errUpdateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	t, err := time.Parse(apiDateFormat, r.PostFormValue(fieldTime))
	if err != nil {
		h.log.Error(errUpdateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	title := r.PostFormValue(fieldTitle)
	category := r.PostFormValue(fieldCategory)

	transaction, err := h.core.UpdateTransaction(userUUID, accountUUID, transactionUUID, t, amount, title, category)
	if err != nil {
		h.log.Error(errUpdateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		h.log.Error(errUpdateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
