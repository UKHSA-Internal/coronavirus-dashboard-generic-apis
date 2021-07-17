package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
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

func FormatPartitionTimestamp(template, timestamp string) (string, error) {

	parsedTimestamp, err := time.Parse(template, timestamp)

	if err != nil {
		return "", err
	}

	return strings.ReplaceAll(parsedTimestamp.Format("2006 1 2"), " ", "_"), nil

} // FormatPartitionTimestamp

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

	return FormatPartitionTimestamp("2006-01-02", results["date"].(string))

} // GetLatestTimestamp

// JSONMarshal is a custom marshal function to serialise JSON payloads
// without escaping HTML characters and converting them to unicode codes.
func JSONMarshal(payload interface{}) ([]byte, error) {

	buffer := &bytes.Buffer{}
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)

	err := encoder.Encode(payload)

	return buffer.Bytes(), err

} // JSONMarshal

func ValidateParam(pattern, value string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(value)
} // ValidateParam
