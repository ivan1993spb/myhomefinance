package handlers

import (
	"net/http"

	"github.com/ivan1993spb/myhomefinance/core"
)

type GetAccountTransactionsByTimeRangeCategoriesHandler struct {
	core *core.Core
}

func (h *GetAccountTransactionsByTimeRangeCategoriesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
