package handlers

import (
	"net/http"

	"github.com/ivan1993spb/myhomefinance/core"
)

type GetAccountTransactionsByTimeRangeHandler struct {
	core *core.Core
}

func (h *GetAccountTransactionsByTimeRangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
