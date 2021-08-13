package api

import (
	"fmt"
	"time"

	"generic_apis/base"
	"generic_apis/insight"
	"github.com/caarlos0/env"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	unit "unit.nginx.org/go"
)

func Run(apiClient *base.Api) {

	var err error

	if err = env.Parse(apiClient); err != nil {
		panic(err)
	}

	addr := fmt.Sprintf(":%s", apiClient.Port)

	// Insight initialisation
	apiClient.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(apiClient.Insight, true)
	apiClient.Routes = UrlPatterns
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

} // Run
