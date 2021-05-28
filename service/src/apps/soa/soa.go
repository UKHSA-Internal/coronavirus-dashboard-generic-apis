package soa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"generic_apis/apps/utils"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/gorilla/mux"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type handler struct {
	db          *db.Config
	traceparent string
}

func (conf *handler) getLatestTimestamp(areaType string) (string, error) {

	category := utils.ReleaseCategories[areaType]

	payload := &db.Payload{
		Query:         timestampQuery,
		Args:          []interface{}{category},
		OperationData: insight.GetOperationData(conf.traceparent),
	}
	results, err := conf.db.FetchRow(payload)
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

} // getLatestTimestamp

func (conf *handler) getPreppedQuery(areaType string) (string, error) {

	areaType = utils.AreaTypes[strings.ToLower(areaType)]
	partition := utils.AreaPartitions[areaType]

	timestamp, err := conf.getLatestTimestamp(areaType)
	if err != nil {
		return "", err
	}

	targetTable := fmt.Sprintf(queryTable, timestamp, partition)
	preppedQuery := fmt.Sprintf(query, targetTable, targetTable)

	return preppedQuery, nil

} // getPreppedQuery

func (conf *handler) fromDatabase(areaType, areaCode string) ([]byte, error) {

	preppedQuery, err := conf.getPreppedQuery(areaType)
	if err != nil {
		return nil, err
	}

	areaCode = strings.ToUpper(areaCode)

	payload := &db.Payload{
		Query:         preppedQuery,
		Args:          []interface{}{areaCode},
		OperationData: insight.GetOperationData(conf.traceparent),
	}
	results, err := conf.db.FetchAll(payload)
	if err != nil {
		return nil, err
	}

	data := make(map[string]interface{})
	for _, item := range results {
		for key, value := range item {
			data[key] = value
		}
	}

	return json.Marshal(data)

} // FromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{}

	return func(w http.ResponseWriter, r *http.Request) {

		var err error

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(insight)
		if err != nil {
			panic(err)
		}
		defer conf.db.CloseConnection()

		pathVars := mux.Vars(r)

		response, err := conf.fromDatabase(pathVars["area_type"], pathVars["area_code"])
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
