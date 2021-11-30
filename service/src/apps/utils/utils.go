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

func (req *GenericRequest) GetLatestPartitionId(areaType string) (string, error) {

	areaType = AreaTypes[strings.ToLower(areaType)]
	category := ReleaseCategories[areaType]

	payload := &db.Payload{
		Query:         processTimestampQuery,
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

} // GetLatestPartitionId

func (req *GenericRequest) GetLatestTimeStamp() (string, error) {

	payload := &db.Payload{
		Query:         timestampQuery,
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

	if _, ok := results["date"]; !ok {
		return "", err
	}

	return FormatPartitionTimestamp("2006-01-02", results["date"].(string))

} // GetLatestPartitionId

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

// ParseAreaPattern converts arrays of area patterns to array of
// human-readable area names.
func ParseAreaPattern(areasPatterns []interface{}) *[]string {
	// Using maps to ensure uniqueness.
	areas := make(map[string]bool, len(areasPatterns))

	for _, area := range areasPatterns {
		switch area {
		case "overview::^K.*$":
			areas["UK"] = true
			continue
		case "nation::^E92000001$":
			areas["England"] = true
			continue
		case "region::^E.*$":
			areas["England regions"] = true
			continue
		case "utla::^E.*$":
			areas["England UTLAs"] = true
			continue
		case "ltla::^E.*$":
			areas["England LTLAs"] = true
			continue
		case "msoa::^E.*$":
			areas["England MSOAs"] = true
			continue
		case "nation::^S92000003$":
			areas["Scotland"] = true
			continue
		case "utla::^S.*$":
		case "ltla::^S.*$":
			areas["Scotland local authorities"] = true
			continue
		case "nation::^N92000002$":
			areas["Northern Ireland"] = true
			continue
		case "utla::^N.*$":
		case "ltla::^N.*$":
			areas["Northern Ireland local authorities"] = true
			continue
		case "nation::^W.*$":
			areas["Wales"] = true
			continue
		case "utla::^W.*$":
		case "ltla::^W.*$":
			areas["Wales local authorities"] = true
			continue
		case "nhsRegion::^.*$":
			areas["All NHS regions"] = true
			continue
		case "nhsTrust::^.*$":
			areas["All NHS trusts"] = true
			continue
		default:
			areas["Unspecified"] = true
			continue
		}
	}

	uniqueAreas := make([]string, len(areas))
	areaInd := 0
	for areaItem := range areas {
		// Deduplicated items leave empty spaces
		// in the map.
		if areaItem != "" {
			uniqueAreas[areaInd] = areaItem
			areaInd++
		}
	}

	return &uniqueAreas

} // ParseAreaPattern
