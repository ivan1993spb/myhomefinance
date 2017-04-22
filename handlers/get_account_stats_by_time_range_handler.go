package handlers

import (
	"net/http"

	"github.com/ivan1993spb/myhomefinance/core"
)

type GetAccountStatsByTimeRangeHandler struct {
	core *core.Core
}

func (h *GetAccountStatsByTimeRangeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
