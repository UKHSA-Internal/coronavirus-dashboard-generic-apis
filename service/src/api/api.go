package api

import (
	"fmt"
	"net/http"
	"time"

	"generic_apis/base"
	"generic_apis/caching"
	"generic_apis/insight"
	"generic_apis/taks_queue"
	"github.com/caarlos0/env"
	"github.com/gorilla/mux"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func Run(apiClient *base.Api) {

	var err error

	if err = env.Parse(apiClient); err != nil {
		panic(err)
	}

	// Insight initialisation
	apiClient.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(apiClient.Insight, true)
	apiClient.Routes = UrlPatterns

	// Redis initialisation
	redisConf := &caching.Config{}
	apiClient.Redis = &caching.RedisClient{
		Client:   redisConf.GetRedisClient(),
		HostName: redisConf.HostName,
	}
	apiClient.Redis.Queue = taks_queue.NewQueue(caching.SetEx(apiClient.Redis), caching.RedisPoolSize)

	defer func() {
		err = apiClient.Redis.Client.Close()
		panic(err)
	}()

	// Initialise the application
	apiClient.Initialize()

	// Launch server
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", apiClient.Port),
		Handler:      apiClient.Router,
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	apiClient.Router = mux.NewRouter()

	if err = srv.ListenAndServe(); err != nil {
		panic(err)
	}

	// Comment for testing
	// This will only run inside the container - needs Nginx Unit
	// to be installed.
	// if err = unit.ListenAndServe(addr, apiClient.Router); err != nil {
	// 	panic(err)
	// }

	// Finalise the app - prepare to exit.
	apiClient.Redis.Queue.FinaliseAndClose(10 * time.Second)

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
