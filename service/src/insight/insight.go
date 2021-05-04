package insight

import (
	"os"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var instrumentationKey = os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY")

func InitialiseInsightClient() appinsights.TelemetryClient {

	var (
		hostName string
		err      error
	)

	insight := appinsights.NewTelemetryClient(instrumentationKey)
	insight.Context().Tags.Cloud().SetRole(getCloudRoleName())
	hostName, err = os.Hostname()
	if err != nil {
		panic(err)
	}
	insight.Context().Tags.Cloud().SetRoleInstance(hostName)

	return insight

} // InitialiseInsightClient
