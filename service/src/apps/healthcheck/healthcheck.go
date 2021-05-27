package healthcheck

import (
	"encoding/json"
	"net/http"

	"generic_apis/db"
	"generic_apis/insight"
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

func Handler(config *db.Config) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{config, ""}

	return func(w http.ResponseWriter, r *http.Request) {

		conf.traceparent = r.Header.Get("traceparent")

		response, err := conf.fromDatabase()
		if err != nil {
			http.Error(w, "failed to retrieve data from the database", http.StatusBadRequest)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
