package utils

import (
	"log"
	"net/url"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type FailedResponse struct {
	HttpCode int
	Response error
	Payload  error
}

func (failure *FailedResponse) Record(insight appinsights.TelemetryClient, url *url.URL) {

	log.Println(failure.Payload.Error())

	exception := appinsights.NewExceptionTelemetry(failure.Payload.Error())

	if failure.HttpCode < 500 {
		exception.SeverityLevel = appinsights.Warning
	} else {
		exception.SeverityLevel = appinsights.Critical
	}

	exception.Properties["url"] = url.Path
	if url.RawQuery != "" {
		exception.Properties["url"] += "?" + url.RawQuery
	}

	insight.TrackException(exception)

} // Record
