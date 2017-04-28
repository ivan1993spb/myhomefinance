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

const URLRouteCreateTransaction = "/user/{" + routeVarUserUUID + "}/account/{" + routeVarAccountUUID + "}/transaction"

const formatURLRouteCreateTransaction = "/user/%s/account/%s/transaction"

func BuildPathCreateTransaction(userUUID, accountUUID uuid.UUID) string {
	return fmt.Sprintf(formatURLRouteCreateTransaction, userUUID, accountUUID)
}

type errCreateTransactionHandler string

func (e errCreateTransactionHandler) Error() string {
	return "error on create transaction handler: " + string(e)
}

type createTransactionHandler struct {
	core *core.Core
	log  *logrus.Logger
}

func NewCreateTransactionHandler(core *core.Core, log *logrus.Logger) http.Handler {
	return &createTransactionHandler{
		core: core,
		log:  log,
	}
}

func (h *createTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	amount, err := strconv.ParseFloat(r.PostFormValue(fieldAmount), 64)
	if err != nil {
		h.log.Error(errCreateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	t, err := time.Parse(apiDateFormat, r.PostFormValue(fieldTime))
	if err != nil {
		h.log.Error(errCreateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	title := r.PostFormValue(fieldTitle)
	category := r.PostFormValue(fieldCategory)

	transaction, err := h.core.CreateTransaction(userUUID, accountUUID, t, amount, title, category)
	if err != nil {
		h.log.Error(errCreateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	// todo set all output to json
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	err = json.NewEncoder(w).Encode(transaction)
	if err != nil {
		h.log.Error(errCreateTransactionHandler(err.Error()))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}
