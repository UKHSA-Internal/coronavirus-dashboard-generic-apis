package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"generic_apis/base"
	"generic_apis/insight"
	"generic_apis/middleware"
	"generic_apis/taks_queue"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type RedisClient struct {
	AzureRedisHost     string `env:"AZURE_REDIS_HOST"`
	AzureRedisPort     string `env:"AZURE_REDIS_PORT"`
	AzureRedisPassword string `env:"AZURE_REDIS_PASSWORD"`
	redisHostName      string
}

const redisPoolSize = 3

func (conf *RedisClient) getRedisClient() *redis.Client {

	if err := env.Parse(conf); err != nil {
		panic(err)
	}

	conf.redisHostName = strings.Split(conf.AzureRedisHost, ".")[0]

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.AzureRedisHost, conf.AzureRedisPort),
		Password: conf.AzureRedisPassword,
		DB:       3,
		PoolSize: redisPoolSize,
	})

	return redisClient

} // getRedisClient

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

	// Redis initialisation
	redisConf := &RedisClient{}
	apiClient.RedisClient = redisConf.getRedisClient()
	apiClient.RedisHostName = redisConf.redisHostName
	defer func() {
		err = apiClient.RedisClient.Close()
		panic(err)
	}()

	// Redis background jobs
	ctx := context.Background()
	setter := func(args interface{}) {
		payload := args.(middleware.SetExPayload)
		apiClient.RedisClient.SetEX(ctx, payload.Key, payload.Value, payload.Duration)
	}

	apiClient.RedisQueue = taks_queue.NewQueue(setter, redisPoolSize)

	// Initialise the application
	apiClient.Initialize()

	// res := apiClient.RedisClient.Ping(context.Background())
	// fmt.Println(res)
	//
	// Uncomment for local testing
	if err = http.ListenAndServe(addr, apiClient.Router); err != nil {
		panic(err)
	}

	// Comment for testing
	// This will only run inside the container - needs Nginx Unit
	// to be installed.
	// if err = unit.ListenAndServe(addr, apiClient.Router); err != nil {
	// 	panic(err)
	// }

	// Finalise the app - prepare to exit.
	apiClient.RedisQueue.FinaliseAndClose(10 * time.Second)

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
