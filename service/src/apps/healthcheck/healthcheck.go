package healthcheck

import (
	"encoding/json"
	"net/http"

	"generic_apis/db"
	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type handler struct {
	db          *db.Config
	traceparent string
}

func (conf *handler) fromDatabase() ([]byte, error) {

	payload := &db.Payload{
		Query:         healthCheckQuery,
		Args:          []interface{}{},
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	_, err := conf.db.FetchAll(payload)
	if err != nil {
		return nil, err
	}

	data := map[string]string{
		"database": "healthy",
	}

	return json.Marshal(data)

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

		response, err := conf.fromDatabase()
		if err != nil {
			http.Error(w, "db connection failed", http.StatusInternalServerError)
			return
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
