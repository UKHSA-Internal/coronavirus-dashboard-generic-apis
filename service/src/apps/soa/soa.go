package soa

import (
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

func (conf *handler) getPreppedQuery(areaType, date string) (string, error) {

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

	dateQuery := ""
	if len(date) != 0 {
		// If date has been requested, use the query param.
		dateQuery = definedDate
	}

	preppedQuery := fmt.Sprintf(query, targetTable, dateQuery)

	return preppedQuery, nil

} // getPreppedQuery

func (conf *handler) fromDatabase(areaType, areaCode, metric, date string) (map[string]interface{}, error) {

	preppedQuery, err := conf.getPreppedQuery(areaType, date)
	if err != nil {
		return nil, err
	}

	areaCode = strings.ToUpper(areaCode)
	queryArgs := []interface{}{areaCode, metric}

	if len(date) != 0 {
		queryArgs = append(queryArgs, date)
	}

	payload := &db.Payload{
		Query:         preppedQuery,
		Args:          queryArgs,
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

	return data, nil

} // fromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{insight: insight}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			err      error
			response map[string]interface{}
			payload  []byte
		)

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(conf.insight)
		if err != nil {
			panic(err)
		}

		pathVars := mux.Vars(r)
		date := r.URL.Query().Get("date")

		response, err = conf.fromDatabase(pathVars["area_type"], pathVars["area_code"], pathVars["metric"], date)
		if err != nil {
			panic(err.Error())
		}

		if len(response) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		if payload, err = utils.JSONMarshal(response); err != nil {
			panic(err)
		} else if _, err = w.Write(payload); err != nil {
			panic(err)
		}

		conf.db.CloseConnection()

	}

} // Handler
