package api

import (
	"fmt"
	"strings"
	"time"

	"generic_apis/base"
	"generic_apis/insight"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
	unit "unit.nginx.org/go"
)

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

	redisConf := &RedisClient{}
	apiClient.RedisClient = redisConf.getRedisClient()
	apiClient.RedisHostName = redisConf.redisHostName
	defer func() {
		err = apiClient.RedisClient.Close()
		panic(err)
	}()

	apiClient.Initialize()

	// res := apiClient.RedisClient.Ping(context.Background())
	// fmt.Println(res)
	//
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
