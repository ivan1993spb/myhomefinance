package outflow

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteOutflowHandlerFunc turns a function with the right signature into a delete outflow handler
type DeleteOutflowHandlerFunc func(DeleteOutflowParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteOutflowHandlerFunc) Handle(params DeleteOutflowParams) middleware.Responder {
	return fn(params)
}

// DeleteOutflowHandler interface for that can handle valid delete outflow params
type DeleteOutflowHandler interface {
	Handle(DeleteOutflowParams) middleware.Responder
}

// NewDeleteOutflow creates a new http.Handler for the delete outflow operation
func NewDeleteOutflow(ctx *middleware.Context, handler DeleteOutflowHandler) *DeleteOutflow {
	return &DeleteOutflow{Context: ctx, Handler: handler}
}

/*DeleteOutflow swagger:route DELETE /outflow outflow deleteOutflow

Deletes outflow document by id

*/
type DeleteOutflow struct {
	Context *middleware.Context
	Handler DeleteOutflowHandler
}

func (o *DeleteOutflow) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewDeleteOutflowParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}