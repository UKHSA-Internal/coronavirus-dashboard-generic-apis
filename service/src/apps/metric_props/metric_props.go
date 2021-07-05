package metric_props

import (
	"net/http"
	"net/url"

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
		preppedQuery string
		args         []interface{}
	)

	if search := params.Get("by"); search == "category" {
		preppedQuery = byCategory
	} else if search == "tag" {
		preppedQuery = byTag
	}

	payload := &db.Payload{
		Query:         preppedQuery,
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

		response, err := conf.fromDatabase(r.URL.Query())
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
