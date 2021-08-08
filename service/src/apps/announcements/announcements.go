package announcements

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

func (conf *handler) fromDatabase(latest bool) ([]db.ResultType, error) {

	var (
		params []interface{}
		query  = allDataQuery
	)

	if latest {
		query = latestDataQuery
	}

	payload := &db.Payload{
		Query:         query,
		Args:          params,
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	res, err := conf.db.FetchAll(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data")
	}

	return res, nil

} // FromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			err     error
			encoded []byte
		)

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(insight)
		if err != nil {
			panic(err)
		}

		isLatest := strings.Contains(r.URL.Path, "/latest")

		response, err := conf.fromDatabase(isLatest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if len(response) == 0 {
			if _, err = w.Write([]byte("[]")); err != nil {
				return
			}
			panic(err)
		}

		encoded, err = utils.JSONMarshal(response)
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(encoded); err != nil {
			panic(err)
		}

		conf.db.CloseConnection()
	}

} // Handler
