package base

import (
	"net/http"

	"generic_apis/apps/healthcheck"
	"generic_apis/apps/utils"
	"generic_apis/db"
	"generic_apis/middleware"
	"generic_apis/taks_queue"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type Api struct {
	Router        *mux.Router
	Routes        *[]utils.RouteEntry
	Database      *db.Config
	Insight       appinsights.TelemetryClient
	Port          string `env:"WEBSITE_PORT"`
	RedisClient   *redis.Client
	RedisHostName string
	RedisQueue    *taks_queue.Queue
}

func (apiClient *Api) Initialize() {

	apiClient.Router = mux.NewRouter()
	apiClient.Router.Use(
		middleware.LogRequest,
		middleware.HeadersMiddleware,
		middleware.PrepareTelemetryMiddleware(apiClient.Insight),
	)

	apiClient.Router.
		Handle(`/generic/healthcheck`, healthcheck.Handler()).
		Name("healthcheck")

	for _, route := range *apiClient.Routes {
		apiClient.Router.
			HandleFunc(
				route.Path,
				middleware.FromCacheOrDB(
					apiClient.RedisClient,
					apiClient.RedisQueue,
					apiClient.RedisHostName,
					apiClient.Insight,
					route.CacheDuration,
					route.Handler,
				),
			).
			Queries().
			Name(route.Name).
			Methods(http.MethodGet)
	}

} // initializeRoutes
