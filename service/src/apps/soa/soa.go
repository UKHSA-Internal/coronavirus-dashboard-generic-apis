package soa

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"generic_apis/apps/utils"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/gorilla/mux"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type handler struct {
	db          *db.Config
	traceparent string
	insight     appinsights.TelemetryClient
}

func (conf *handler) getPreppedQuery(areaType string) (string, error) {

	areaType = utils.AreaTypes[strings.ToLower(areaType)]
	partition := utils.AreaPartitions[areaType]

	req := &utils.GenericRequest{
		Traceparent: conf.traceparent,
		Insight:     conf.insight,
	}
	timestamp, err := req.GetLatestTimestamp(areaType)
	if err != nil {
		return "", err
	}

	targetTable := fmt.Sprintf(queryTable, timestamp, partition)
	preppedQuery := fmt.Sprintf(query, targetTable, targetTable)

	return preppedQuery, nil

} // getPreppedQuery

func (conf *handler) fromDatabase(areaType, areaCode, metric string) ([]byte, error) {

	preppedQuery, err := conf.getPreppedQuery(areaType)
	if err != nil {
		return nil, err
	}

	areaCode = strings.ToUpper(areaCode)

	payload := &db.Payload{
		Query:         preppedQuery,
		Args:          []interface{}{areaCode, metric},
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

	conf := &handler{insight: insight}

	return func(w http.ResponseWriter, r *http.Request) {

		var err error

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(conf.insight)
		if err != nil {
			panic(err)
		}
		defer conf.db.CloseConnection()

		pathVars := mux.Vars(r)

		response, err := conf.fromDatabase(pathVars["area_type"], pathVars["area_code"], pathVars["metric"])
		if err != nil {
			panic(err.Error())
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
