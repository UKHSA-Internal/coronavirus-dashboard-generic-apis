package metric_search

import (
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

func (conf *handler) fromDatabase(params url.Values) ([]byte, error) {

	var (
		filters string
		args    []interface{}
		counter = 0
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

	payload := &db.Payload{
		Query:         fmt.Sprintf(query, filters),
		Args:          args,
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	results, err := conf.db.FetchAll(payload)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return []byte{'[', ']'}, nil
	}

	return utils.JSONMarshal(results)

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

		response, err := conf.fromDatabase(r.URL.Query())
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

		conf.db.CloseConnection()

	}

} // queryByCode
