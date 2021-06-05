package code

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

func (conf *handler) fromDatabase(areaType, code string) ([]byte, error) {

	var (
		ok     bool
		params []interface{}
		query  string
	)

	code = strings.ReplaceAll(strings.ToUpper(code), " ", "")
	areaType = strings.ToLower(areaType)

	if areaType, ok = utils.AreaTypes[areaType]; !ok {
		return nil, fmt.Errorf("invalid area type")
	}

	if areaType == utils.AreaTypes["postcode"] {
		params = []interface{}{utils.AreaTypes["msoa"], code}
		query = postcodeQuery
	} else {
		params = []interface{}{areaType, code}
		query = areaQuery
	}

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

	if len(results) > 0 && areaType == utils.AreaTypes["postcode"] {
		postcode := results[0]["postcode"].(string)
		data["postcode"] = postcode
		data["trimmedPostcode"] = strings.ReplaceAll(postcode, " ", "")
	}

	return utils.JSONMarshal(data)

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

		pathVars := mux.Vars(r)

		response, err := conf.fromDatabase(pathVars["area_type"], pathVars["area_code"])
		if err != nil {
			http.Error(w, "failed to retrieve data", http.StatusBadRequest)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
