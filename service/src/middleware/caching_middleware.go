package middleware

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type HandlerFunc func(appinsights.TelemetryClient) func(http.ResponseWriter, *http.Request)

var ctx = context.Background()

func FromCacheOrDB(redisClient *redis.Client, redisHostName string, insight appinsights.TelemetryClient,
	cacheDuration time.Duration, handler HandlerFunc) http.HandlerFunc {

	handlerFunc := handler(insight)

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			payload     string
			err         error
			startTime   time.Time
			endTime     time.Time
			redisAction string
		)

		if cacheDuration == 0 {
			handlerFunc(w, r)
			w.Header().Set("X-CACHE-HIT", "0")
			return
		} else {
			redisAction = "GET"
			startTime = time.Now()
			payload, err = redisClient.Get(ctx, r.RequestURI).Result()
			endTime = time.Now()
		}

		if err == redis.Nil {

			w.Header().Set("X-CACHE-HIT", "0")

			rec := httptest.NewRecorder()
			handlerFunc(rec, r)
			w.WriteHeader(rec.Result().StatusCode)

			_, err = io.Copy(w, rec.Result().Body)
			if err != nil {
				panic(err)
			}

			redisAction = "SET"
			startTime = time.Now()
			redisClient.SetEX(ctx, r.RequestURI, rec.Result().Body, cacheDuration)
			endTime = time.Now()

		} else if err != nil {

			panic(err)

		} else {

			res := []byte(payload)
			_, err = w.Write(res)
			if err != nil {
				panic(err)
			}

			w.Header().Set("X-CACHE-HIT", "1")

		}

		dependency := appinsights.NewRemoteDependencyTelemetry(
			redisHostName,
			"Redis",
			"redis",
			err == nil,
		)
		dependency.Data = r.RequestURI
		dependency.Properties["action"] = redisAction
		dependency.MarkTime(startTime, endTime)

	}

} // FromCacheOrDB
