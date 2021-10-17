package base

import (
	"fmt"
	"net/http"
	"strings"

	"generic_apis/apps/healthcheck"
	"generic_apis/apps/utils"
	"generic_apis/db"
	"generic_apis/middleware"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type Api struct {
	Router   *mux.Router
	Routes   *[]utils.RouteEntry
	Database *db.Config
	Insight  appinsights.TelemetryClient
	Port     string `env:"WEBSITE_PORT"`
}

type RedisClient struct {
	AzureRedisHost     string `env:"AZURE_REDIS_HOST"`
	AzureRedisPort     string `env:"AZURE_REDIS_PORT"`
	AzureRedisPassword string `env:"AZURE_REDIS_PASSWORD"`
	redisHostName      string
}

func (conf *RedisClient) getRedisClient() *redis.Client {

	if err := env.Parse(conf); err != nil {
		panic(err)
	}

	conf.redisHostName = strings.Split(conf.AzureRedisHost, ".")[0]

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.AzureRedisHost, conf.AzureRedisPort),
		Password: conf.AzureRedisPassword,
		DB:       3,
		PoolSize: 10,
		PoolFIFO: false,
	})

	return redisClient

} // getRedisClient

func (apiClient *Api) Initialize() {

	redisConf := &RedisClient{}
	redisClient := redisConf.getRedisClient()

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
					redisClient,
					redisConf.redisHostName,
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
