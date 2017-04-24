package handlers

import (
	"net/http"

	"github.com/Sirupsen/logrus"

	"github.com/ivan1993spb/myhomefinance/core"
)

type getAccountTransactionsByTimeRangeCategoriesHandler struct {
	core *core.Core
	log  *logrus.Logger
}

func NewGetAccountTransactionsByTimeRangeCategoriesHandler(core *core.Core, log *logrus.Logger) http.Handler {
	return &getAccountTransactionsByTimeRangeCategoriesHandler{
		core: core,
		log:  log,
	}
}

func (h *getAccountTransactionsByTimeRangeCategoriesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
