package middleware

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	"generic_apis/caching"
	"github.com/go-redis/redis/v8"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type HandlerFunc func(appinsights.TelemetryClient) func(http.ResponseWriter, *http.Request)

var ctx = context.Background()

func FromCacheOrDB(redisCli *caching.RedisClient, insight appinsights.TelemetryClient, cacheDuration time.Duration,
	handler HandlerFunc) http.HandlerFunc {

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
			w.Header().Set("X-CACHE-HIT", "-1")
			return
		} else {
			redisAction = "GET"
			startTime = time.Now()
			payload, err = redisCli.Client.Get(ctx, r.RequestURI).Result()
			endTime = time.Now()
		}

		if err == redis.Nil {

			w.Header().Set("X-CACHE-HIT", "0")

			rec := httptest.NewRecorder()
			handlerFunc(rec, r)
			statusCode := rec.Result().StatusCode
			w.WriteHeader(statusCode)

			data := bytes.NewBuffer(nil)
			_, err = io.Copy(data, rec.Result().Body)
			_, _ = w.Write(data.Bytes())
			if err != nil {
				panic(err)
			}
			if statusCode >= 400 {
				return
			}

			redisAction = "SET"
			startTime = time.Now()
			setExPayload := &caching.SetExPayload{
				Key:      r.RequestURI,
				Value:    data.Bytes(),
				Duration: cacheDuration,
			}
			redisCli.Queue.Push(setExPayload)
			endTime = time.Now()

		} else if err != nil {

			panic(err)

		} else {

			res := []byte(payload)
			w.Header().Set("X-CACHE-HIT", "1")
			_, err = w.Write(res)
			if err != nil {
				panic(err)
			}

		}

		dependency := appinsights.NewRemoteDependencyTelemetry(
			redisCli.HostName,
			"Redis",
			"caching",
			err == nil,
		)
		dependency.Data = r.RequestURI
		dependency.Properties["action"] = redisAction
		dependency.MarkTime(startTime, endTime)

	}

} // FromCacheOrDB
