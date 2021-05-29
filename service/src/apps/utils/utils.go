package utils

import (
	"fmt"
	"strings"
	"time"

	"generic_apis/db"
	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type GenericRequest struct {
	Traceparent string
	Insight     appinsights.TelemetryClient
}

func (req *GenericRequest) GetLatestTimestamp(areaType string) (string, error) {

	areaType = AreaTypes[strings.ToLower(areaType)]
	category := ReleaseCategories[areaType]

	payload := &db.Payload{
		Query:         timestampQuery,
		Args:          []interface{}{category},
		OperationData: insight.GetOperationData(req.Traceparent),
	}

	database, err := db.Connect(req.Insight)
	if err != nil {
		return "", err
	}
	defer database.CloseConnection()

	results, err := database.FetchRow(payload)
	if err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no valid timestamp for '%s'", areaType)
	}

	if _, ok := results["date"]; !ok {
		return "", err
	}

	date, _ := time.Parse("2006-01-02", results["date"].(string))

	result := strings.ReplaceAll(date.Format("2006 1 2"), " ", "_")

	return result, nil

} // GetLatestTimestamp
