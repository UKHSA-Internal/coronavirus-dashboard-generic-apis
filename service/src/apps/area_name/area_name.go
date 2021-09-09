package area_name

import (
	"fmt"
	"net/http"
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

func (conf *handler) fromDatabase(areaType, areaName string) (map[string]string, error) {

	var (
		ok     bool
		params []interface{}
		query  string
	)

	areaType = strings.ToLower(areaType)

	if areaType, ok = utils.AreaTypes[areaType]; !ok {
		return nil, fmt.Errorf("invalid area type")
	}

	params = []interface{}{areaType, areaName}
	query = areaQuery

	payload := &db.Payload{
		Query:         query,
		Args:          params,
		OperationData: insight.GetOperationData(conf.traceparent),
	}
	results, err := conf.db.FetchAll(payload)
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	for _, item := range results {
		data[item["areaType"].(string)+"Name"] = item["areaName"].(string)
		data[item["areaType"].(string)] = item["areaCode"].(string)
	}

	return data, nil

} // fromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			err         error
			response    map[string]string
			jsonPayload []byte
		)

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(insight)
		if err != nil {
			panic(err)
		}

		pathVars := mux.Vars(r)

		response, err = conf.fromDatabase(pathVars["area_type"], pathVars["area_name"])
		if err != nil {
			http.Error(w, "failed to retrieve data", http.StatusBadRequest)
			return
		}

		if len(response) == 0 {
			http.NotFound(w, r)
			return
		}

		jsonPayload, err = utils.JSONMarshal(response)
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(jsonPayload); err != nil {
			panic(err)
		}

		conf.db.CloseConnection()

	}

} // Handler
