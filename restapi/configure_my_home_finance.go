package restapi

import (
	"crypto/tls"
	"log"
	"net/http"
	"os"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/jessevdk/go-flags"

	"database/sql"

	"github.com/ivan1993spb/myhomefinance/mappers"
	"github.com/ivan1993spb/myhomefinance/restapi/operations"
	"github.com/ivan1993spb/myhomefinance/restapi/operations/additionally"
	"github.com/ivan1993spb/myhomefinance/restapi/operations/inflow"
	"github.com/ivan1993spb/myhomefinance/restapi/operations/notes"
	"github.com/ivan1993spb/myhomefinance/restapi/operations/outflow"
	"github.com/ivan1993spb/myhomefinance/restapi/operations/transactions"
)

// This file is safe to edit. Once it exists it will not be overwritten

type ServerOptions struct {
	DBFile flags.Filename `short:"f" long:"db-file" description:"sqlite3 db file name" default:"myhomefinance.db"`
}

var serverOptions ServerOptions

func configureFlags(api *operations.MyHomeFinanceAPI) {
	api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{
		{
			ShortDescription: "MyHomeFinance Options",
			LongDescription:  "",
			Options:          &serverOptions,
		},
	}
}

func configureAPI(api *operations.MyHomeFinanceAPI) http.Handler {
	api.ServeError = errors.ServeError

	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags)
	api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	var (
		noteMapper mappers.NoteMapper
	)

	api.InflowDeleteInflowHandler = inflow.DeleteInflowHandlerFunc(func(params inflow.DeleteInflowParams) middleware.Responder {
		return middleware.NotImplemented("operation inflow.DeleteInflow has not yet been implemented")
	})
	api.NotesDeleteNotesHandler = notes.DeleteNotesHandlerFunc(func(params notes.DeleteNotesParams) middleware.Responder {
		return middleware.NotImplemented("operation notes.DeleteNotes has not yet been implemented")
	})
	api.OutflowDeleteOutflowHandler = outflow.DeleteOutflowHandlerFunc(func(params outflow.DeleteOutflowParams) middleware.Responder {
		return middleware.NotImplemented("operation outflow.DeleteOutflow has not yet been implemented")
	})
	api.AdditionallyGetDestinationsGrepHandler = additionally.GetDestinationsGrepHandlerFunc(func(params additionally.GetDestinationsGrepParams) middleware.Responder {
		return middleware.NotImplemented("operation additionally.GetDestinationsGrep has not yet been implemented")
	})
	api.InflowGetInflowHandler = inflow.GetInflowHandlerFunc(func(params inflow.GetInflowParams) middleware.Responder {
		return middleware.NotImplemented("operation inflow.GetInflow has not yet been implemented")
	})
	api.InflowGetInflowDateFromDateToHandler = inflow.GetInflowDateFromDateToHandlerFunc(func(params inflow.GetInflowDateFromDateToParams) middleware.Responder {
		return middleware.NotImplemented("operation inflow.GetInflowDateFromDateTo has not yet been implemented")
	})
	api.InflowGetInflowDateFromDateToGrepHandler = inflow.GetInflowDateFromDateToGrepHandlerFunc(func(params inflow.GetInflowDateFromDateToGrepParams) middleware.Responder {
		return middleware.NotImplemented("operation inflow.GetInflowDateFromDateToGrep has not yet been implemented")
	})
	api.AdditionallyGetMetricUnitsHandler = additionally.GetMetricUnitsHandlerFunc(func(params additionally.GetMetricUnitsParams) middleware.Responder {
		return middleware.NotImplemented("operation additionally.GetMetricUnits has not yet been implemented")
	})
	api.NotesGetNotesHandler = notes.GetNotesHandlerFunc(func(params notes.GetNotesParams) middleware.Responder {
		note, err := noteMapper.GetNoteById(params.ID)
		if err != nil {
			if err == sql.ErrNoRows {
				return notes.NewGetNotesNotFound()
			}
			return notes.NewGetNotesServiceUnavailable()
		}
		return notes.NewGetNotesOK().WithPayload(note)
	})
	api.NotesGetNotesDateFromDateToHandler = notes.GetNotesDateFromDateToHandlerFunc(func(params notes.GetNotesDateFromDateToParams) middleware.Responder {
		return middleware.NotImplemented("operation notes.GetNotesDateFromDateTo has not yet been implemented")
	})
	api.NotesGetNotesDateFromDateToGrepHandler = notes.GetNotesDateFromDateToGrepHandlerFunc(func(params notes.GetNotesDateFromDateToGrepParams) middleware.Responder {
		return middleware.NotImplemented("operation notes.GetNotesDateFromDateToGrep has not yet been implemented")
	})
	api.OutflowGetOutflowHandler = outflow.GetOutflowHandlerFunc(func(params outflow.GetOutflowParams) middleware.Responder {
		return middleware.NotImplemented("operation outflow.GetOutflow has not yet been implemented")
	})
	api.OutflowGetOutflowDateFromDateToHandler = outflow.GetOutflowDateFromDateToHandlerFunc(func(params outflow.GetOutflowDateFromDateToParams) middleware.Responder {
		return middleware.NotImplemented("operation outflow.GetOutflowDateFromDateTo has not yet been implemented")
	})
	api.OutflowGetOutflowDateFromDateToGrepHandler = outflow.GetOutflowDateFromDateToGrepHandlerFunc(func(params outflow.GetOutflowDateFromDateToGrepParams) middleware.Responder {
		return middleware.NotImplemented("operation outflow.GetOutflowDateFromDateToGrep has not yet been implemented")
	})
	api.AdditionallyGetSourcesGrepHandler = additionally.GetSourcesGrepHandlerFunc(func(params additionally.GetSourcesGrepParams) middleware.Responder {
		return middleware.NotImplemented("operation additionally.GetSourcesGrep has not yet been implemented")
	})
	api.AdditionallyGetStatisticsHandler = additionally.GetStatisticsHandlerFunc(func(params additionally.GetStatisticsParams) middleware.Responder {
		return middleware.NotImplemented("operation additionally.GetStatistics has not yet been implemented")
	})
	api.AdditionallyGetTargetsGrepHandler = additionally.GetTargetsGrepHandlerFunc(func(params additionally.GetTargetsGrepParams) middleware.Responder {
		return middleware.NotImplemented("operation additionally.GetTargetsGrep has not yet been implemented")
	})
	api.TransactionsGetTransactionsHandler = transactions.GetTransactionsHandlerFunc(func(params transactions.GetTransactionsParams) middleware.Responder {
		return middleware.NotImplemented("operation transactions.GetTransactions has not yet been implemented")
	})
	api.TransactionsGetTransactionsDateFromDateToHandler = transactions.GetTransactionsDateFromDateToHandlerFunc(func(params transactions.GetTransactionsDateFromDateToParams) middleware.Responder {
		return middleware.NotImplemented("operation transactions.GetTransactionsDateFromDateTo has not yet been implemented")
	})
	api.TransactionsGetTransactionsDateFromDateToGrepHandler = transactions.GetTransactionsDateFromDateToGrepHandlerFunc(func(params transactions.GetTransactionsDateFromDateToGrepParams) middleware.Responder {
		return middleware.NotImplemented("operation transactions.GetTransactionsDateFromDateToGrep has not yet been implemented")
	})
	api.InflowPostInflowHandler = inflow.PostInflowHandlerFunc(func(params inflow.PostInflowParams) middleware.Responder {
		return middleware.NotImplemented("operation inflow.PostInflow has not yet been implemented")
	})
	api.NotesPostNotesHandler = notes.PostNotesHandlerFunc(func(params notes.PostNotesParams) middleware.Responder {
		return middleware.NotImplemented("operation notes.PostNotes has not yet been implemented")
	})
	api.OutflowPostOutflowHandler = outflow.PostOutflowHandlerFunc(func(params outflow.PostOutflowParams) middleware.Responder {
		return middleware.NotImplemented("operation outflow.PostOutflow has not yet been implemented")
	})
	api.InflowPutInflowHandler = inflow.PutInflowHandlerFunc(func(params inflow.PutInflowParams) middleware.Responder {
		return middleware.NotImplemented("operation inflow.PutInflow has not yet been implemented")
	})
	api.NotesPutNotesHandler = notes.PutNotesHandlerFunc(func(params notes.PutNotesParams) middleware.Responder {
		return middleware.NotImplemented("operation notes.PutNotes has not yet been implemented")
	})
	api.OutflowPutOutflowHandler = outflow.PutOutflowHandlerFunc(func(params outflow.PutOutflowParams) middleware.Responder {
		return middleware.NotImplemented("operation outflow.PutOutflow has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
