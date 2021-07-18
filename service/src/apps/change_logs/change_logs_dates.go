package change_logs

import (
	"net/http"

	"generic_apis/apps/utils"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func (conf *handler) getDatesFromDatabase() ([]db.ResultType, error) {

	payload := &db.Payload{
		Query:         recordMonths,
		Args:          []interface{}{},
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	res, err := conf.db.FetchAll(payload)

	return res, err

} // getDatesFromDatabase

func DatesHandler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

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

		response, err := conf.getDatesFromDatabase()
		if err != nil {
			http.Error(w, "failed to retrieve data", http.StatusBadRequest)
		}

		if len(response) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
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

} // queryByCode
