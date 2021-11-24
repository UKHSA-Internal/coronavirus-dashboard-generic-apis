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

type failedResponse struct {
	httpCode int
	response error
	payload  error
}

func (conf *handler) fromDatabase(params url.Values) ([]byte, *failedResponse) {

	var (
		filters string
		args    []interface{}
		err     error
		counter = 0
		failure = &failedResponse{}
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
			failure.httpCode = http.StatusBadRequest
			failure.response = errors.New("`exact` flag may only be used alongside the `search` parameter")
			failure.payload = failure.response
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
		failure.httpCode = http.StatusInternalServerError
		failure.response = errors.New("failed to retrieve the requested data for a valid query")
		failure.payload = err
		return nil, failure
	}

	if len(results) == 0 {
		return []byte("[]"), nil
	}

	var response []byte
	if response, err = utils.JSONMarshal(results); err != nil {
		failure.httpCode = http.StatusInternalServerError
		failure.response = errors.New("failed to generate JSON payload")
		failure.payload = err
		return nil, failure
	} else {
		return response, nil
	}

} // fromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			failure *failedResponse
			err     error
		)

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(insight)
		if err != nil {
			panic(err)
		}

		response, failure := conf.fromDatabase(r.URL.Query())
		if failure != nil {
			w.WriteHeader(failure.httpCode)
			http.Error(w, failure.response.Error(), failure.httpCode)
			panic(err)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

		conf.db.CloseConnection()

	}

} // Handler
