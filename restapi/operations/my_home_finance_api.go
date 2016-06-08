package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ivan1993spb/myhomefinance/restapi/operations/additionally"
	"github.com/ivan1993spb/myhomefinance/restapi/operations/inflow"
	"github.com/ivan1993spb/myhomefinance/restapi/operations/notes"
	"github.com/ivan1993spb/myhomefinance/restapi/operations/outflow"
	"github.com/ivan1993spb/myhomefinance/restapi/operations/transactions"
)

// NewMyHomeFinanceAPI creates a new MyHomeFinance instance
func NewMyHomeFinanceAPI(spec *loads.Document) *MyHomeFinanceAPI {
	o := &MyHomeFinanceAPI{
		spec:            spec,
		handlers:        make(map[string]map[string]http.Handler),
		formats:         strfmt.Default,
		defaultConsumes: "application/json",
		defaultProduces: "application/json",
		ServerShutdown:  func() {},
	}

	return o
}

/*MyHomeFinanceAPI This is small finance program named myHomeFinance
 */
type MyHomeFinanceAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// InflowDeleteInflowHandler sets the operation handler for the delete inflow operation
	InflowDeleteInflowHandler inflow.DeleteInflowHandler
	// NotesDeleteNotesHandler sets the operation handler for the delete notes operation
	NotesDeleteNotesHandler notes.DeleteNotesHandler
	// OutflowDeleteOutflowHandler sets the operation handler for the delete outflow operation
	OutflowDeleteOutflowHandler outflow.DeleteOutflowHandler
	// AdditionallyGetDestinationsGrepHandler sets the operation handler for the get destinations grep operation
	AdditionallyGetDestinationsGrepHandler additionally.GetDestinationsGrepHandler
	// InflowGetInflowHandler sets the operation handler for the get inflow operation
	InflowGetInflowHandler inflow.GetInflowHandler
	// InflowGetInflowDateFromDateToHandler sets the operation handler for the get inflow date from date to operation
	InflowGetInflowDateFromDateToHandler inflow.GetInflowDateFromDateToHandler
	// InflowGetInflowDateFromDateToGrepHandler sets the operation handler for the get inflow date from date to grep operation
	InflowGetInflowDateFromDateToGrepHandler inflow.GetInflowDateFromDateToGrepHandler
	// AdditionallyGetMetricUnitsHandler sets the operation handler for the get metric units operation
	AdditionallyGetMetricUnitsHandler additionally.GetMetricUnitsHandler
	// NotesGetNotesHandler sets the operation handler for the get notes operation
	NotesGetNotesHandler notes.GetNotesHandler
	// NotesGetNotesDateFromDateToHandler sets the operation handler for the get notes date from date to operation
	NotesGetNotesDateFromDateToHandler notes.GetNotesDateFromDateToHandler
	// NotesGetNotesDateFromDateToGrepHandler sets the operation handler for the get notes date from date to grep operation
	NotesGetNotesDateFromDateToGrepHandler notes.GetNotesDateFromDateToGrepHandler
	// OutflowGetOutflowHandler sets the operation handler for the get outflow operation
	OutflowGetOutflowHandler outflow.GetOutflowHandler
	// OutflowGetOutflowDateFromDateToHandler sets the operation handler for the get outflow date from date to operation
	OutflowGetOutflowDateFromDateToHandler outflow.GetOutflowDateFromDateToHandler
	// OutflowGetOutflowDateFromDateToGrepHandler sets the operation handler for the get outflow date from date to grep operation
	OutflowGetOutflowDateFromDateToGrepHandler outflow.GetOutflowDateFromDateToGrepHandler
	// AdditionallyGetSourcesGrepHandler sets the operation handler for the get sources grep operation
	AdditionallyGetSourcesGrepHandler additionally.GetSourcesGrepHandler
	// AdditionallyGetStatisticsHandler sets the operation handler for the get statistics operation
	AdditionallyGetStatisticsHandler additionally.GetStatisticsHandler
	// AdditionallyGetTargetsGrepHandler sets the operation handler for the get targets grep operation
	AdditionallyGetTargetsGrepHandler additionally.GetTargetsGrepHandler
	// TransactionsGetTransactionsHandler sets the operation handler for the get transactions operation
	TransactionsGetTransactionsHandler transactions.GetTransactionsHandler
	// TransactionsGetTransactionsDateFromDateToHandler sets the operation handler for the get transactions date from date to operation
	TransactionsGetTransactionsDateFromDateToHandler transactions.GetTransactionsDateFromDateToHandler
	// TransactionsGetTransactionsDateFromDateToGrepHandler sets the operation handler for the get transactions date from date to grep operation
	TransactionsGetTransactionsDateFromDateToGrepHandler transactions.GetTransactionsDateFromDateToGrepHandler
	// InflowPostInflowHandler sets the operation handler for the post inflow operation
	InflowPostInflowHandler inflow.PostInflowHandler
	// NotesPostNotesHandler sets the operation handler for the post notes operation
	NotesPostNotesHandler notes.PostNotesHandler
	// OutflowPostOutflowHandler sets the operation handler for the post outflow operation
	OutflowPostOutflowHandler outflow.PostOutflowHandler
	// InflowPutInflowHandler sets the operation handler for the put inflow operation
	InflowPutInflowHandler inflow.PutInflowHandler
	// NotesPutNotesHandler sets the operation handler for the put notes operation
	NotesPutNotesHandler notes.PutNotesHandler
	// OutflowPutOutflowHandler sets the operation handler for the put outflow operation
	OutflowPutOutflowHandler outflow.PutOutflowHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *MyHomeFinanceAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *MyHomeFinanceAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// DefaultProduces returns the default produces media type
func (o *MyHomeFinanceAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *MyHomeFinanceAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *MyHomeFinanceAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *MyHomeFinanceAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the MyHomeFinanceAPI
func (o *MyHomeFinanceAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.InflowDeleteInflowHandler == nil {
		unregistered = append(unregistered, "inflow.DeleteInflowHandler")
	}

	if o.NotesDeleteNotesHandler == nil {
		unregistered = append(unregistered, "notes.DeleteNotesHandler")
	}

	if o.OutflowDeleteOutflowHandler == nil {
		unregistered = append(unregistered, "outflow.DeleteOutflowHandler")
	}

	if o.AdditionallyGetDestinationsGrepHandler == nil {
		unregistered = append(unregistered, "additionally.GetDestinationsGrepHandler")
	}

	if o.InflowGetInflowHandler == nil {
		unregistered = append(unregistered, "inflow.GetInflowHandler")
	}

	if o.InflowGetInflowDateFromDateToHandler == nil {
		unregistered = append(unregistered, "inflow.GetInflowDateFromDateToHandler")
	}

	if o.InflowGetInflowDateFromDateToGrepHandler == nil {
		unregistered = append(unregistered, "inflow.GetInflowDateFromDateToGrepHandler")
	}

	if o.AdditionallyGetMetricUnitsHandler == nil {
		unregistered = append(unregistered, "additionally.GetMetricUnitsHandler")
	}

	if o.NotesGetNotesHandler == nil {
		unregistered = append(unregistered, "notes.GetNotesHandler")
	}

	if o.NotesGetNotesDateFromDateToHandler == nil {
		unregistered = append(unregistered, "notes.GetNotesDateFromDateToHandler")
	}

	if o.NotesGetNotesDateFromDateToGrepHandler == nil {
		unregistered = append(unregistered, "notes.GetNotesDateFromDateToGrepHandler")
	}

	if o.OutflowGetOutflowHandler == nil {
		unregistered = append(unregistered, "outflow.GetOutflowHandler")
	}

	if o.OutflowGetOutflowDateFromDateToHandler == nil {
		unregistered = append(unregistered, "outflow.GetOutflowDateFromDateToHandler")
	}

	if o.OutflowGetOutflowDateFromDateToGrepHandler == nil {
		unregistered = append(unregistered, "outflow.GetOutflowDateFromDateToGrepHandler")
	}

	if o.AdditionallyGetSourcesGrepHandler == nil {
		unregistered = append(unregistered, "additionally.GetSourcesGrepHandler")
	}

	if o.AdditionallyGetStatisticsHandler == nil {
		unregistered = append(unregistered, "additionally.GetStatisticsHandler")
	}

	if o.AdditionallyGetTargetsGrepHandler == nil {
		unregistered = append(unregistered, "additionally.GetTargetsGrepHandler")
	}

	if o.TransactionsGetTransactionsHandler == nil {
		unregistered = append(unregistered, "transactions.GetTransactionsHandler")
	}

	if o.TransactionsGetTransactionsDateFromDateToHandler == nil {
		unregistered = append(unregistered, "transactions.GetTransactionsDateFromDateToHandler")
	}

	if o.TransactionsGetTransactionsDateFromDateToGrepHandler == nil {
		unregistered = append(unregistered, "transactions.GetTransactionsDateFromDateToGrepHandler")
	}

	if o.InflowPostInflowHandler == nil {
		unregistered = append(unregistered, "inflow.PostInflowHandler")
	}

	if o.NotesPostNotesHandler == nil {
		unregistered = append(unregistered, "notes.PostNotesHandler")
	}

	if o.OutflowPostOutflowHandler == nil {
		unregistered = append(unregistered, "outflow.PostOutflowHandler")
	}

	if o.InflowPutInflowHandler == nil {
		unregistered = append(unregistered, "inflow.PutInflowHandler")
	}

	if o.NotesPutNotesHandler == nil {
		unregistered = append(unregistered, "notes.PutNotesHandler")
	}

	if o.OutflowPutOutflowHandler == nil {
		unregistered = append(unregistered, "outflow.PutOutflowHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *MyHomeFinanceAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *MyHomeFinanceAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *MyHomeFinanceAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *MyHomeFinanceAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *MyHomeFinanceAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

func (o *MyHomeFinanceAPI) initHandlerCache() {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/inflow"] = inflow.NewDeleteInflow(o.context, o.InflowDeleteInflowHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/notes"] = notes.NewDeleteNotes(o.context, o.NotesDeleteNotesHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers[strings.ToUpper("DELETE")] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/outflow"] = outflow.NewDeleteOutflow(o.context, o.OutflowDeleteOutflowHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/destinations/grep"] = additionally.NewGetDestinationsGrep(o.context, o.AdditionallyGetDestinationsGrepHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/inflow"] = inflow.NewGetInflow(o.context, o.InflowGetInflowHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/inflow/{date_from}_{date_to}"] = inflow.NewGetInflowDateFromDateTo(o.context, o.InflowGetInflowDateFromDateToHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/inflow/{date_from}_{date_to}/grep"] = inflow.NewGetInflowDateFromDateToGrep(o.context, o.InflowGetInflowDateFromDateToGrepHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/metric-units"] = additionally.NewGetMetricUnits(o.context, o.AdditionallyGetMetricUnitsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/notes"] = notes.NewGetNotes(o.context, o.NotesGetNotesHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/notes/{date_from}_{date_to}"] = notes.NewGetNotesDateFromDateTo(o.context, o.NotesGetNotesDateFromDateToHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/notes/{date_from}_{date_to}/grep"] = notes.NewGetNotesDateFromDateToGrep(o.context, o.NotesGetNotesDateFromDateToGrepHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/outflow"] = outflow.NewGetOutflow(o.context, o.OutflowGetOutflowHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/outflow/{date_from}_{date_to}"] = outflow.NewGetOutflowDateFromDateTo(o.context, o.OutflowGetOutflowDateFromDateToHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/outflow/{date_from}_{date_to}/grep"] = outflow.NewGetOutflowDateFromDateToGrep(o.context, o.OutflowGetOutflowDateFromDateToGrepHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/sources/grep"] = additionally.NewGetSourcesGrep(o.context, o.AdditionallyGetSourcesGrepHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/statistics"] = additionally.NewGetStatistics(o.context, o.AdditionallyGetStatisticsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/targets/grep"] = additionally.NewGetTargetsGrep(o.context, o.AdditionallyGetTargetsGrepHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/transactions"] = transactions.NewGetTransactions(o.context, o.TransactionsGetTransactionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/transactions/{date_from}_{date_to}"] = transactions.NewGetTransactionsDateFromDateTo(o.context, o.TransactionsGetTransactionsDateFromDateToHandler)

	if o.handlers["GET"] == nil {
		o.handlers[strings.ToUpper("GET")] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/transactions/{date_from}_{date_to}/grep"] = transactions.NewGetTransactionsDateFromDateToGrep(o.context, o.TransactionsGetTransactionsDateFromDateToGrepHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/inflow"] = inflow.NewPostInflow(o.context, o.InflowPostInflowHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/notes"] = notes.NewPostNotes(o.context, o.NotesPostNotesHandler)

	if o.handlers["POST"] == nil {
		o.handlers[strings.ToUpper("POST")] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/outflow"] = outflow.NewPostOutflow(o.context, o.OutflowPostOutflowHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/inflow"] = inflow.NewPutInflow(o.context, o.InflowPutInflowHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/notes"] = notes.NewPutNotes(o.context, o.NotesPutNotesHandler)

	if o.handlers["PUT"] == nil {
		o.handlers[strings.ToUpper("PUT")] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/outflow"] = outflow.NewPutOutflow(o.context, o.OutflowPutOutflowHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *MyHomeFinanceAPI) Serve(builder middleware.Builder) http.Handler {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}

	return o.context.APIHandler(builder)
}
