package additionally

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetStatisticsHandlerFunc turns a function with the right signature into a get statistics handler
type GetStatisticsHandlerFunc func(GetStatisticsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStatisticsHandlerFunc) Handle(params GetStatisticsParams) middleware.Responder {
	return fn(params)
}

// GetStatisticsHandler interface for that can handle valid get statistics params
type GetStatisticsHandler interface {
	Handle(GetStatisticsParams) middleware.Responder
}

// NewGetStatistics creates a new http.Handler for the get statistics operation
func NewGetStatistics(ctx *middleware.Context, handler GetStatisticsHandler) *GetStatistics {
	return &GetStatistics{Context: ctx, Handler: handler}
}

/*GetStatistics swagger:route GET /statistics additionally getStatistics

Returns list of metrix

*/
type GetStatistics struct {
	Context *middleware.Context
	Handler GetStatisticsHandler
}

func (o *GetStatistics) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetStatisticsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}