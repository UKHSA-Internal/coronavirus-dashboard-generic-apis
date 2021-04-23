package soa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"generic_apis/db"
	"github.com/gorilla/mux"
)

const query = `
SELECT 
	area_code AS "areaCode", 
	area_name AS "areaName",
	area_type AS "areaType",
	DATE(date)::TEXT AS "date",
	(payload -> 'rollingSum') AS "rollingSum",
	(payload -> 'rollingRate') AS "rollingRate",
	(payload -> 'change') AS "change",
	(payload -> 'direction') AS "direction",
	(payload -> 'changePercentage') AS "changePercentage"
FROM %s AS ts
	JOIN covid19.area_reference AS ar ON ts.area_id = ar.id
WHERE area_code = $1
  AND date = ( SELECT MAX(date) FROM %s ) 
`

const queryTable = "covid19.time_series_p%s_%s"

const timestampQuery = `
SELECT DATE(MAX(timestamp))::TEXT AS date
FROM covid19.release_reference AS rr 
	JOIN covid19.release_category AS rc ON rc.release_id = rr.id
WHERE released IS TRUE
  AND process_name = $1
`

var areaTypes = map[string]string{
	"postcode":  "postcode",
	"msoa":      "msoa",
	"nhstrust":  "nhsTrust",
	"nhsregion": "nhsRegion",
	"utla":      "utla",
	"ltla":      "ltla",
	"region":    "region",
	"nation":    "nation",
	"overview":  "overview",
}

var releaseCategories = map[string]string{
	"msoa":      "MSOA",
	"nhsTrust":  "MAIN",
	"nhsRegion": "MAIN",
	"utla":      "MAIN",
	"ltla":      "MAIN",
	"region":    "MAIN",
	"nation":    "MAIN",
	"overview":  "MAIN",
}

var areaPartitions = map[string]string{
	"msoa":      "msoa",
	"nhsTrust":  "nhstrust",
	"nhsRegion": "other",
	"utla":      "utla",
	"ltla":      "ltla",
	"region":    "other",
	"nation":    "other",
	"overview":  "other",
}

func getLatestTimestamp(areaType string) (string, error) {

	category := releaseCategories[areaType]
	results, err := db.Query(timestampQuery, category)
	if err != nil {
		return "", err
	}

	if len(results) == 0 {
		return "", fmt.Errorf("no valid timestamp for '%s'", areaType)
	}

	if _, ok := results[0]["date"]; !ok {
		return "", err
	}

	date, _ := time.Parse("2006-01-02", results[0]["date"].(string))

	result := strings.ReplaceAll(date.Format("2006 1 2"), " ", "_")

	return result, nil

} // getLatestTimestamp

func fromDatabase(areaType, areaCode string) ([]byte, error) {

	areaType = areaTypes[strings.ToLower(areaType)]
	partition := areaPartitions[areaType]

	timestamp, err := getLatestTimestamp(areaType)
	if err != nil {
		return nil, err
	}

	targetTable := fmt.Sprintf(queryTable, timestamp, partition)
	preppedQuery := fmt.Sprintf(query, targetTable, targetTable)
	areaCode = strings.ToUpper(areaCode)

	results, err := db.Query(preppedQuery, areaCode)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	for _, item := range results {
		for key, value := range item {
			data[key] = value
		}
	}

	if jsonString, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		return jsonString, nil
	}

} // FromDatabase

func SoaQuery(w http.ResponseWriter, r *http.Request) {

	pathVars := mux.Vars(r)

	var (
		areaType string
		areaCode string
		ok       bool
	)

	if areaType, ok = pathVars["area_type"]; !ok {
		panic("no category")
	} else if areaCode, ok = pathVars["area_code"]; !ok {
		panic("no code")
	}

	response, err := fromDatabase(areaType, areaCode)
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(response); err != nil {
		panic(err)
	}

} // queryByCode
