package metric_availability

import (
	"errors"
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
	if areaCode != "" && strings.ToLower(areaType) != "msoa" {
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
		date, err = req.GetLatestPartitionId(areaType)
	}

	if err != nil {
		return "", err
	}

	targetTable := fmt.Sprintf(queryTable, date, partition)
	preppedQuery := fmt.Sprintf(query, targetTable, areaCodeQuery, targetTable)

	return preppedQuery, nil

} // getPreppedQuery

func (conf *handler) fromDatabase(areaType, areaCode, date string) ([]db.ResultType, *utils.FailedResponse) {

	failure := &utils.FailedResponse{}

	areaTypeLower := strings.ToLower(areaType)
	areaCodeLower := strings.ToLower(areaCode)

	// MSOAs are only available for England areas.
	// Metrics associated with them are consistent across all MSOAs,
	// which means that they do not need to be individually verified.
	if areaTypeLower == "msoa" && (areaCodeLower != "" && !strings.HasPrefix(areaCodeLower, "e")) {
		failure.Response = errors.New(
			"metric availability queries for MSOAs must either be generic " +
				"(no area code) or use an England area code")
		failure.Payload = failure.Response
		failure.HttpCode = http.StatusBadRequest

		return nil, failure
	}

	preppedQuery, err := conf.getPreppedQuery(areaType, areaCode, date)
	if err != nil {
		failure.Response = errors.New("unable to format query")
		failure.Payload = err
		failure.HttpCode = http.StatusBadRequest
		return nil, failure
	}

	queryArgs := []interface{}{areaType}

	// If defined, add area code to query args.
	if areaCode != "" && strings.ToLower(areaType) != "msoa" {
		queryArgs = append(queryArgs, strings.ToUpper(areaCode))
	}

	payload := &db.Payload{
		Query:         preppedQuery,
		Args:          queryArgs,
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	results, resErr := conf.db.FetchAll(payload)
	if resErr != nil {
		failure.Response = errors.New("internal server error")
		failure.Payload = resErr
		failure.HttpCode = http.StatusBadRequest
		return nil, failure
	}

	return results, nil

} // fromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{insight: insight}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			err         error
			response    []db.ResultType
			failure     *utils.FailedResponse
			jsonPayload []byte
		)

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(conf.insight)
		if err != nil {
			panic(err)
		}

		pathVars := mux.Vars(r)
		date := r.URL.Query().Get("date")

		response, failure = conf.fromDatabase(pathVars["area_type"], pathVars["area_code"], date)
		if failure != nil {
			failure.Record(insight, r.URL)
			http.Error(w, failure.Response.Error(), failure.HttpCode)
			return
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
