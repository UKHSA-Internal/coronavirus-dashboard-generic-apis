package caching

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type SetExPayload struct {
	Key       string
	Value     []byte
	Duration  time.Duration
	Telemetry *TelemetryPayload
}

type TelemetryPayload struct {
	Insight       appinsights.TelemetryClient
	RedisHostName string
	Key           string
	Action        string
	Start         time.Time
	End           time.Time
	Err           error
}

type Config struct {
	AzureRedisHost     string `env:"AZURE_REDIS_HOST"`
	AzureRedisPort     string `env:"AZURE_REDIS_PORT"`
	AzureRedisPassword string `env:"AZURE_REDIS_PASSWORD"`
	HostName           string
}

type RedisClient struct {
	Client   *redis.Client
	HostName string
}

const (
	RedisMinClients   = 25
	RedisDB           = 3
	hostUrlDelimiter  = "."
	redisAddrTemplate = "%s:%s"
	SetCache          = "SET"
	GetCache          = "GET"
)

var ctx = context.Background()

func (payload *TelemetryPayload) Push() {

	dependency := appinsights.NewRemoteDependencyTelemetry(
		payload.RedisHostName,
		"Redis",
		"caching",
		payload.Err == nil,
	)
	dependency.Data = payload.Key
	dependency.Properties["action"] = payload.Action
	dependency.MarkTime(payload.Start, payload.End)
	payload.Insight.Track(dependency)

} // pushTelemetry

func (conf *Config) GetRedisClient() *redis.Client {

	if err := env.Parse(conf); err != nil {
		log.Fatal(err)
	}

	conf.HostName = strings.Split(conf.AzureRedisHost, hostUrlDelimiter)[0]

	redisOpts := &redis.Options{
		Addr:         fmt.Sprintf(redisAddrTemplate, conf.AzureRedisHost, conf.AzureRedisPort),
		Password:     conf.AzureRedisPassword,
		DB:           RedisDB,
		MinIdleConns: RedisMinClients,
	}

	redisClient := redis.NewClient(redisOpts)

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Panic(err)
	}

	return redisClient

} // getRedisClient
