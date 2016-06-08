package outflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetOutflowHandlerFunc turns a function with the right signature into a get outflow handler
type GetOutflowHandlerFunc func(GetOutflowParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetOutflowHandlerFunc) Handle(params GetOutflowParams) middleware.Responder {
	return fn(params)
}

// GetOutflowHandler interface for that can handle valid get outflow params
type GetOutflowHandler interface {
	Handle(GetOutflowParams) middleware.Responder
}

// NewGetOutflow creates a new http.Handler for the get outflow operation
func NewGetOutflow(ctx *middleware.Context, handler GetOutflowHandler) *GetOutflow {
	return &GetOutflow{Context: ctx, Handler: handler}
}

/*GetOutflow swagger:route GET /outflow outflow getOutflow

GetOutflow get outflow API

*/
type GetOutflow struct {
	Context *middleware.Context
	Handler GetOutflowHandler
}

func (o *GetOutflow) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetOutflowParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
