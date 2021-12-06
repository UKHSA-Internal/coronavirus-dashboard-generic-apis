package metric_search

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"generic_apis/apps/utils"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type handler struct {
	db          *db.Config
	traceparent string
}

func (conf *handler) fromDatabase(params url.Values) ([]byte, *utils.FailedResponse) {

	var (
		filters string
		args    []interface{}
		err     error
		counter = 0
		failure = &utils.FailedResponse{}
	)

	if search := params.Get("search"); search != "" {
		counter += 1
		args = append(args, fmt.Sprintf(metricTemplate, strings.ToLower(search)))
		filters += fmt.Sprintf(searchFilter, counter, counter)
	}

	if category := params.Get("category"); category != "" {
		counter += 1
		args = append(args, strings.ToLower(category))
		filters += fmt.Sprintf(categoryFilter, counter)
	}

	if tags := params.Get("tags"); tags != "" {
		counter += 1
		args = append(args, strings.Split(strings.ToLower(tags), ","))
		filters += fmt.Sprintf(tagsFilter, counter)
	}

	if isExact := params.Get("exact"); isExact != "" && isExact != "0" {
		if search := params.Get("search"); search == "" {
			failure.HttpCode = http.StatusBadRequest
			failure.Response = errors.New("`exact` flag may only be used alongside the `search` parameter")
			failure.Payload = failure.Response
			return nil, failure
		} else {
			counter = 1
			filters = fmt.Sprintf(searchExactFilter, counter)
		}

	}

	payload := &db.Payload{
		Query:         fmt.Sprintf(query, filters),
		Args:          args,
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	var results []db.ResultType
	results, err = conf.db.FetchAll(payload)
	if err != nil {
		failure.HttpCode = http.StatusInternalServerError
		failure.Response = errors.New("failed to retrieve the requested data for a valid query")
		failure.Payload = err

		return nil, failure
	}

	if len(results) == 0 {
		failure.HttpCode = http.StatusNotFound
		failure.Response = errors.New("not found")
		failure.Payload = errors.New("not found")

		return nil, failure
	}

	var response []byte
	response, err = utils.JSONMarshal(results)
	if err != nil {
		failure.HttpCode = http.StatusInternalServerError
		failure.Response = errors.New("failed to generate JSON payload")
		failure.Payload = err

		return nil, failure
	}

	return response, nil

} // fromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			failure *utils.FailedResponse
			err     error
		)
		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(insight)
		if err != nil {
			panic(err)
		}

		response, failure := conf.fromDatabase(r.URL.Query())
		if failure != nil {
			failure.Record(insight, r.URL)
			http.Error(w, failure.Response.Error(), failure.HttpCode)
			return
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

		conf.db.CloseConnection()

	}

} // Handler
