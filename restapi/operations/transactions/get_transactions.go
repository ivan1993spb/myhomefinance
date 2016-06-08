package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetTransactionsHandlerFunc turns a function with the right signature into a get transactions handler
type GetTransactionsHandlerFunc func(GetTransactionsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTransactionsHandlerFunc) Handle(params GetTransactionsParams) middleware.Responder {
	return fn(params)
}

// GetTransactionsHandler interface for that can handle valid get transactions params
type GetTransactionsHandler interface {
	Handle(GetTransactionsParams) middleware.Responder
}

// NewGetTransactions creates a new http.Handler for the get transactions operation
func NewGetTransactions(ctx *middleware.Context, handler GetTransactionsHandler) *GetTransactions {
	return &GetTransactions{Context: ctx, Handler: handler}
}

/*GetTransactions swagger:route GET /transactions transactions getTransactions

GetTransactions get transactions API

*/
type GetTransactions struct {
	Context *middleware.Context
	Handler GetTransactionsHandler
}

func (o *GetTransactions) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetTransactionsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}