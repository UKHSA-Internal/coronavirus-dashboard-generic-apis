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

func FromCacheOrDB(redisCli *caching.RedisClient, insight appinsights.TelemetryClient, cacheDuration time.Duration,
	handler HandlerFunc) http.HandlerFunc {

	handlerFunc := handler(insight)

	return func(w http.ResponseWriter, r *http.Request) {

		var payload string
		ctx := context.Background()

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

			if statusCode > 200 {
				telemetry.Push()
				return
			}

			if statusCode >= http.StatusBadRequest {
				return
			}

			setExPayload := &caching.SetExPayload{
				Key:       r.RequestURI,
				Value:     data.Bytes(),
				Duration:  cacheDuration,
				Telemetry: telemetry,
			}
			redisCli.Queue.Push(setExPayload)

		} else if telemetry.Err != nil {

			telemetry.Push()
			panic(telemetry.Err)

		} else {

			res := []byte(payload)
			w.Header().Set(cacheHeader, cacheHit)
			defer telemetry.Push()

			if _, telemetry.Err = w.Write(res); telemetry.Err != nil {
				panic(telemetry.Err)
			}

		}

	}

} // FromCacheOrDB
