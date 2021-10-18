package middleware

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"time"

	"generic_apis/taks_queue"
	"github.com/go-redis/redis/v8"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type HandlerFunc func(appinsights.TelemetryClient) func(http.ResponseWriter, *http.Request)

type SetExPayload struct {
	Key      string
	Value    []byte
	Duration time.Duration
}

var ctx = context.Background()

func FromCacheOrDB(redisClient *redis.Client, redisQueue *taks_queue.Queue, redisHostName string,
	insight appinsights.TelemetryClient, cacheDuration time.Duration, handler HandlerFunc) http.HandlerFunc {

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
			payload, err = redisClient.Get(ctx, r.RequestURI).Result()
			endTime = time.Now()
		}

		if err == redis.Nil {

			w.Header().Set("X-CACHE-HIT", "0")

			rec := httptest.NewRecorder()
			handlerFunc(rec, r)
			w.WriteHeader(rec.Result().StatusCode)

			data := bytes.NewBuffer(nil)
			_, err = io.Copy(data, rec.Result().Body)
			_, _ = w.Write(data.Bytes())
			if err != nil {
				panic(err)
			}

			redisAction = "SET"
			startTime = time.Now()
			payload := SetExPayload{
				Key:      r.RequestURI,
				Value:    data.Bytes(),
				Duration: cacheDuration,
			}
			redisQueue.Push(payload)
			redisClient.SetEX(ctx, r.RequestURI, data.Bytes(), cacheDuration)
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
