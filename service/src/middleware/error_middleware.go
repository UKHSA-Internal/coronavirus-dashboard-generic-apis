package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var isDev = os.Getenv("IS_DEV") == "1"

func PrepareHandlerErrorMiddleware(insightClient appinsights.TelemetryClient) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			uri := r.URL.Path
			status := http.StatusOK
			success := true

			defer func() {

				r := recover()

				if r != nil {
					var err error

					switch t := r.(type) {
					case string:
						err = errors.New(t)
					case error:
						err = t
					default:
						err = errors.New("unknown error")
					}

					success = false
					exception := appinsights.NewExceptionTelemetry(err)
					exception.SeverityLevel = appinsights.Warning
					exception.Properties["url"] = uri
					insightClient.Track(exception)

					status = http.StatusInternalServerError
					http.Error(w, "Internal server error", status)

					if isDev {
						panic(err)
					}

				}

			}()

			start := time.Now()
			traceparent := insight.GetOperationData(r.Header.Get("traceparent"))
			request := appinsights.NewRequestTelemetry("GET", uri, 0, fmt.Sprintf("%d", status))

			r.Header.Set("traceparent", traceparent.TraceParent)
			request.Id = traceparent.ParentId
			request.Tags.Operation().SetId(traceparent.OperationId)
			request.Tags.Operation().SetParentId(traceparent.OperationId)
			request.Tags.Operation().SetName(fmt.Sprintf("GET %s", uri))

			next.ServeHTTP(w, r)
			w.Header().Set("traceparent", traceparent.TraceParent)
			request.Success = success
			request.MarkTime(start, time.Now())
			insightClient.Track(request)

		})

	}

} // PrepareHandlerErrorMiddleware
