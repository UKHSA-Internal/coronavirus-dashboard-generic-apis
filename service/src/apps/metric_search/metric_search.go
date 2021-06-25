package metric_search

import (
	"fmt"
	"net/http"
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

func (conf *handler) fromDatabase(params map[string]string) ([]byte, error) {

	var (
		searchToken = fmt.Sprintf(metricTemplate, strings.ToLower(params["search"]))
		args        = []interface{}{searchToken}
	)

	payload := &db.Payload{
		Query:         query,
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
		defer conf.db.CloseConnection()

		// pathVars := mux.Vars(r)
		// fmt.Println(fmt.Sprintf("%v", pathVars))

		searchParam := r.URL.Query().Get("search")
		if searchParam == "" {
			panic("invalid search param")
		}

		params := map[string]string{"search": searchParam}

		response, err := conf.fromDatabase(params)
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
