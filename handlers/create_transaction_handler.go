package handlers

import (
	"net/http"

	"github.com/ivan1993spb/myhomefinance/core"
)

type CreateTransactionHandler struct {
	core *core.Core
}

func (h *CreateTransactionHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}
