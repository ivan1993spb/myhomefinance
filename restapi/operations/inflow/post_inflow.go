package inflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostInflowHandlerFunc turns a function with the right signature into a post inflow handler
type PostInflowHandlerFunc func(PostInflowParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostInflowHandlerFunc) Handle(params PostInflowParams) middleware.Responder {
	return fn(params)
}

// PostInflowHandler interface for that can handle valid post inflow params
type PostInflowHandler interface {
	Handle(PostInflowParams) middleware.Responder
}

// NewPostInflow creates a new http.Handler for the post inflow operation
func NewPostInflow(ctx *middleware.Context, handler PostInflowHandler) *PostInflow {
	return &PostInflow{Context: ctx, Handler: handler}
}

/*PostInflow swagger:route POST /inflow inflow postInflow

PostInflow post inflow API

*/
type PostInflow struct {
	Context *middleware.Context
	Handler PostInflowHandler
}

func (o *PostInflow) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewPostInflowParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
