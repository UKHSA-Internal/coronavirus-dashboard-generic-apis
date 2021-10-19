package base

import (
	"net/http"

	"generic_apis/apps/healthcheck"
	"generic_apis/apps/utils"
	"generic_apis/caching"
	"generic_apis/db"
	"generic_apis/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type Api struct {
	Router   *mux.Router
	Routes   *[]utils.RouteEntry
	Database *db.Config
	Insight  appinsights.TelemetryClient
	Port     string `env:"WEBSITE_PORT"`
	Redis    *caching.RedisClient
}

const (
	openApiFilePath = "/opt/app/assets"
	openApiUri      = "/generic/openapi.json"
	heathCheckPath  = "/generic/healthcheck"
	healthCheckName = "healthcheck"
)

func (apiClient *Api) Initialize() {

	// Setting the middleware
	apiClient.Router.Use(
		handlers.ProxyHeaders,
		middleware.LogRequest,
		middleware.HeadersMiddleware,
		middleware.PrepareTelemetryMiddleware(apiClient.Insight),
	)

	// Health check
	apiClient.Router.
		HandleFunc(heathCheckPath, healthcheck.Handler()).
		Name(healthCheckName)

	// Static files
	fs := http.FileServer(http.Dir(openApiFilePath))
	apiClient.Router.Handle(openApiUri, fs)

	// API routes
	for _, route := range *apiClient.Routes {
		apiClient.Router.
			HandleFunc(
				route.Path,
				middleware.FromCacheOrDB(
					apiClient.Redis,
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
