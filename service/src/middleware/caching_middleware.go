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

const (
	cacheHeader = "X-CACHE-HIT"
	cacheHit    = "1"
	cacheMiss   = "0"
	uncached    = "-1"
	bypassCache = 0
)

var ctx = context.Background()

func FromCacheOrDB(redisCli *caching.RedisClient, insight appinsights.TelemetryClient, cacheDuration time.Duration,
	handler HandlerFunc) http.HandlerFunc {

	handlerFunc := handler(insight)

	return func(w http.ResponseWriter, r *http.Request) {

		var payload string

		telemetry := &caching.TelemetryPayload{
			Insight:       insight,
			Key:           r.RequestURI,
			RedisHostName: redisCli.HostName,
		}

		if cacheDuration == bypassCache {
			handlerFunc(w, r)
			w.Header().Set(cacheHeader, uncached)
			return
		} else {
			telemetry.Action = caching.GetCache
			telemetry.Start = time.Now()
			payload, telemetry.Err = redisCli.Client.Get(ctx, r.RequestURI).Result()
			telemetry.End = time.Now()
		}

		if telemetry.Err == redis.Nil {

			telemetry.Err = nil
			telemetry.Action = caching.SetCache

			w.Header().Set(cacheHeader, cacheMiss)

			rec := httptest.NewRecorder()
			handlerFunc(rec, r)
			statusCode := rec.Result().StatusCode
			w.WriteHeader(statusCode)

			data := bytes.NewBuffer(nil)

			_, telemetry.Err = io.Copy(data, rec.Result().Body)
			if telemetry.Err != nil {
				telemetry.Push()
				panic(telemetry.Err)
			}

			_, telemetry.Err = w.Write(data.Bytes())
			if telemetry.Err != nil {
				telemetry.Push()
				panic(telemetry.Err)
			}

			if statusCode >= http.StatusBadRequest {
				return
			}

			telemetry.Start = time.Now()
			redisCli.Client.SetEX(ctx, r.RequestURI, data.Bytes(), cacheDuration)
			telemetry.End = time.Now()
			telemetry.Push()

		} else if telemetry.Err != nil {

			telemetry.Push()
			panic(telemetry.Err)

		} else {

			res := []byte(payload)
			w.Header().Set(cacheHeader, cacheHit)

			if _, telemetry.Err = w.Write(res); telemetry.Err != nil {
				telemetry.Push()
				panic(telemetry.Err)
			}

		}

	}

} // FromCacheOrDB
