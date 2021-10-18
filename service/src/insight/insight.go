package insight

import (
	"os"
	"time"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

var instrumentationKey = os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY")

func InitialiseInsightClient() appinsights.TelemetryClient {

	insightConf := appinsights.NewTelemetryConfiguration(instrumentationKey)
	insightConf.MaxBatchSize = 1024
	insightConf.MaxBatchInterval = 2 * time.Second

	insight := appinsights.NewTelemetryClientFromConfig(insightConf)
	insight.Context().Tags.Cloud().SetRole(GetCloudRoleName())
	insight.Context().Tags.Cloud().SetRoleInstance(GetCloudRoleInstance())

	return insight

} // InitialiseInsightClient
