package api

import (
	"net/http"
	"time"

	"generic_apis/apps/areaByType"
	"generic_apis/apps/code"
	"generic_apis/apps/pageArea"
	"generic_apis/apps/soa"
	"generic_apis/db"
	"generic_apis/insight"
	"generic_apis/middleware"
	"github.com/gorilla/mux"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	"unit.nginx.org/go"
)

type (
	Api struct {
		Router   *mux.Router
		database *db.Config
		Insight  appinsights.TelemetryClient
	}

	routeEntry struct {
		name    string
		path    string
		handler func(*db.Config) func(http.ResponseWriter, *http.Request)
	}
)

var getRoutes = []routeEntry{
	{
		"code",
		`/generic/code/{area_type:[a-zA-Z]{4,10}}/{area_code:[a-zA-Z0-9+%\s]{3,12}}`,
		code.Handler,
	},
	{
		"soa",
		`/generic/soa/{area_type:[a-zA-Z]{4,10}}/{area_code:[a-zA-Z0-9]{3,10}}`,
		soa.Handler,
	},
	{
		"area",
		`/generic/area/{area_type:[a-zA-Z]{4,10}}`,
		areaByType.Handler,
	},
	{
		"page_areas",
		`/generic/page_areas/{page:[a-zA-Z]{3,10}}`,
		pageArea.Handler,
	},
	{
		"page_areas_with_type",
		`/generic/page_areas/{page:[a-zA-Z]{3,10}}/{area_type:[a-zA-Z]{5,12}}`,
		pageArea.Handler,
	},
} // routes

func (apiClient *Api) Run(addr string) {

	// Insight initialisation
	apiClient.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(apiClient.Insight, true)

	// DB initialisation
	var err error
	apiClient.database, err = db.Connect(apiClient.Insight)
	if err != nil {
		panic(err)
	}

	apiClient.Initialize()

	// Uncomment for local testing
	// if err = http.ListenAndServe(addr, apiClient.Router); err != nil {
	// 	panic(err)
	// }

	// Comment for testing
	// This will only run inside the container - needs Nginx Unit
	// to be installed.
	if err = unit.ListenAndServe(addr, apiClient.Router); err != nil {
		panic(err)
	}

	select {
	case <-apiClient.Insight.Channel().Close(10 * time.Second):
		// Ten second timeout for retries.
		// All telemetry should have been submitted
		// successfully and we can proceed to exiting.

	case <-time.After(30 * time.Second):
		// Thirty second absolute timeout.  This covers any
		// previous telemetry submission that may not have
		// completed before Close was called.
		// There are a number of reasons we could have
		// reached here.  Telemetry submission has likely
		// failed somewhere.
	}

	defer apiClient.database.CloseConnection()

} // Run

func (apiClient *Api) Initialize() {

	telemetryMiddleware := middleware.PrepareTelemetryMiddleware(apiClient.Insight)

	apiClient.Router = mux.NewRouter()
	apiClient.Router.Use(middleware.HeadersMiddleware)
	apiClient.Router.Use(telemetryMiddleware)

	for _, route := range getRoutes {
		apiClient.Router.
			HandleFunc(route.path, route.handler(apiClient.database)).
			Name(route.name).
			Methods("GET")
	}

} // initializeRoutes
