package insight

import (
	"os"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var instrumentationKey = os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY")

func InitialiseInsightClient() appinsights.TelemetryClient {

	insight := appinsights.NewTelemetryClient(instrumentationKey)
	insight.Context().Tags.Cloud().SetRole(GetCloudRoleName())
	insight.Context().Tags.Cloud().SetRoleInstance(GetCloudRoleInstance())

	return insight

} // InitialiseInsightClient
