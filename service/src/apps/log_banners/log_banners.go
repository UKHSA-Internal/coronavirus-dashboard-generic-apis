package log_banners

import (
	"fmt"
	"net/http"
	"strconv"
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
}

func (conf *handler) fromDatabase(urlParams *map[string]string) ([]db.ResultType, error) {

	var (
		params []interface{}
		query  = mainQuery
		pcount = 0
	)

	for key, value := range *urlParams {
		pcount += 1
		query = strings.ReplaceAll(query, fmt.Sprintf(`{%s_token}`, key), strconv.Itoa(pcount))
		params = append(params, value)
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

} // fromDatabase

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

		pathVars := mux.Vars(r)

		response, err := conf.fromDatabase(&pathVars)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		if len(response) == 0 {
			if _, err = w.Write([]byte("[]")); err != nil {
				return
			}
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

} // Handler
