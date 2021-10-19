package caching

import (
	"context"
	"fmt"
	"strings"
	"time"

	"generic_apis/taks_queue"
	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
)

type SetExPayload struct {
	Key      string
	Value    []byte
	Duration time.Duration
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
	Queue    *taks_queue.Queue
}

const (
	RedisMinClients   = 5
	RedisDB           = 3
	hostUrlDelimiter  = "."
	redisAddrTemplate = "%s:%s"
)

func (conf *Config) GetRedisClient() *redis.Client {

	if err := env.Parse(conf); err != nil {
		panic(err)
	}

	conf.HostName = strings.Split(conf.AzureRedisHost, hostUrlDelimiter)[0]

	redisOpts := &redis.Options{
		Addr:         fmt.Sprintf(redisAddrTemplate, conf.AzureRedisHost, conf.AzureRedisPort),
		Password:     conf.AzureRedisPassword,
		DB:           RedisDB,
		MinIdleConns: RedisMinClients,
	}

	redisClient := redis.NewClient(redisOpts)

	return redisClient

} // getRedisClient

func SetEx(redis *RedisClient) func(args interface{}) {

	ctx := context.Background()

	return func(args interface{}) {

		payload := args.(*SetExPayload)

		redis.Client.SetEX(ctx, payload.Key, payload.Value, payload.Duration)

	}

} // SetEx
