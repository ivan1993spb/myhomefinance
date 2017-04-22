package handlers

import (
	"net/http"

	"github.com/ivan1993spb/myhomefinance/core"
)

type UpdateTransactionHandler struct {
	core *core.Core
}

func (h *UpdateTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
