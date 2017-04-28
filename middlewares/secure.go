package middlewares

import (
	"net/http"

	"github.com/urfave/negroni"
)

var Secure negroni.HandlerFunc = func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(rw,r)
}
