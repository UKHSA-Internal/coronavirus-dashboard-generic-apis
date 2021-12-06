package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var isDev = os.Getenv("IS_DEV") == "1"

func PrepareTelemetryMiddleware(insightClient appinsights.TelemetryClient) func(next http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			uri := r.URL.Path
			if r.URL.RawQuery != "" {
				uri += "?" + r.URL.RawQuery
			}
			status := http.StatusOK
			success := true

			defer func() {

				log.SetFlags(log.LstdFlags)

				rec := recover()

				if rec != nil {
					var err error

					switch t := rec.(type) {
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
					insightClient.TrackException(exception)

					status = http.StatusInternalServerError
					log.Println(err)
					http.Error(w, "Internal server error", status)

					if isDev {
						panic(err)
					}

					return
				}

			}()

			start := time.Now()
			optData := insight.GetOperationData(r.Header.Get("traceparent"))
			request := appinsights.NewRequestTelemetry("GET", uri, 0, fmt.Sprintf("%d", status))

			r.Header.Set("traceparent", optData.TraceParent)
			request.Id = optData.ParentId
			request.Tags.Operation().SetId(optData.OperationId)
			request.Tags.Operation().SetParentId(optData.OperationId)
			request.Tags.Operation().SetName(fmt.Sprintf("GET %s", uri))
			request.Tags.Cloud().SetRole(optData.CloudRoleName)
			request.Tags.Cloud().SetRoleInstance(optData.CloudRoleInstance)

			next.ServeHTTP(w, r)
			w.Header().Set("traceparent", optData.TraceParent)
			request.Success = success
			request.MarkTime(start, time.Now())
			insightClient.Track(request)

		})

	}

} // PrepareTelemetryMiddleware
