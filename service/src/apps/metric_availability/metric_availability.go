package metric_availability

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

func (conf *handler) getPreppedQuery(areaType, areaCode, date string) (string, error) {

	var err error

	areaType = utils.AreaTypes[strings.ToLower(areaType)]
	partition := utils.AreaPartitions[areaType]

	// Decide whether to include
	// area code in the query.
	areaCodeQuery := ""
	if areaCode != "" {
		areaCodeQuery = areaCodeFilter
	}

	req := &utils.GenericRequest{
		Traceparent: conf.traceparent,
		Insight:     conf.insight,
	}

	// Determine the partition date based on
	// whether the request is against the
	// latest data or the archives.
	if date != "" {
		date, err = utils.FormatPartitionTimestamp("2006-01-02", date)
	} else {
		date, err = req.GetLatestTimestamp(areaType)
	}

	if err != nil {
		return "", err
	}

	targetTable := fmt.Sprintf(queryTable, date, partition)
	preppedQuery := fmt.Sprintf(query, targetTable, areaCodeQuery, targetTable)

	return preppedQuery, nil

} // getPreppedQuery

func (conf *handler) fromDatabase(areaType, areaCode, date string) ([]db.ResultType, error) {

	preppedQuery, err := conf.getPreppedQuery(areaType, areaCode, date)
	if err != nil {
		return nil, err
	}

	queryArgs := []interface{}{areaType}

	// If defined, add area code to query args.
	if areaCode != "" {
		queryArgs = append(queryArgs, strings.ToUpper(areaCode))
	}

	payload := &db.Payload{
		Query:         preppedQuery,
		Args:          queryArgs,
		OperationData: insight.GetOperationData(conf.traceparent),
	}
	return conf.db.FetchAll(payload)

} // fromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{insight: insight}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			err         error
			response    []db.ResultType
			jsonPayload []byte
		)

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(conf.insight)
		if err != nil {
			panic(err)
		}

		pathVars := mux.Vars(r)
		date := r.URL.Query().Get("date")

		response, err = conf.fromDatabase(pathVars["area_type"], pathVars["area_code"], date)
		if err != nil {
			panic(err.Error())
		}

		if len(response) == 0 {
			http.NotFound(w, r)
			return
		}

		jsonPayload, err = utils.JSONMarshal(response)
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(jsonPayload); err != nil {
			panic(err)
		}

		conf.db.CloseConnection()

	}

} // Handler
