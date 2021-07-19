package area_by_type

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

func (conf *handler) fromDatabase(areaType string) ([]byte, error) {

	var ok bool
	if areaType, ok = utils.AreaTypes[strings.ToLower(areaType)]; !ok {
		return nil, fmt.Errorf("invalid area type")
	}

	payload := &db.Payload{
		Query:         query,
		Args:          []interface{}{areaType},
		OperationData: insight.GetOperationData(conf.traceparent),
	}
	results, err := conf.db.FetchAll(payload)
	if err != nil {
		return nil, err
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

		pathVars := mux.Vars(r)

		response, err := conf.fromDatabase(pathVars["area_type"])
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

		conf.db.CloseConnection()

	}

} // queryByCode
