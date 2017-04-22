package handlers

import (
	"net/http"

	"github.com/ivan1993spb/myhomefinance/core"
)

type DeleteTransactionHandler struct {
	core *core.Core
}

func (h *DeleteTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
