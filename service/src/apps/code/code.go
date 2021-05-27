package code

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"generic_apis/apps/utils"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/gorilla/mux"
)

type handler struct {
	db          *db.Config
	traceparent string
}

func (conf *handler) fromDatabase(areaType, search string) ([]byte, error) {

	var (
		ok     bool
		params []interface{}
		query  string
	)

	search = strings.ReplaceAll(strings.ToUpper(search), " ", "")
	areaType = strings.ToLower(areaType)

	if areaType, ok = utils.AreaTypes[areaType]; !ok {
		return nil, fmt.Errorf("invalid area type")
	}

	if areaType == utils.AreaTypes["postcode"] {
		params = []interface{}{utils.AreaTypes["msoa"], search}
		query = postcodeQuery
	} else {
		params = []interface{}{areaType, search}
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

	return json.Marshal(data)

} // FromDatabase

func Handler(config *db.Config) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{config, ""}

	return func(w http.ResponseWriter, r *http.Request) {

		conf.traceparent = r.Header.Get("traceparent")

		pathVars := mux.Vars(r)

		var (
			category string
			code     string
			ok       bool
		)

		if category, ok = pathVars["area_type"]; !ok {
			http.Error(w, "area type not defined", http.StatusBadRequest)
			return
		} else if code, ok = pathVars["area_code"]; !ok {
			http.Error(w, "area code not defined", http.StatusBadRequest)
		}

		response, err := conf.fromDatabase(category, code)
		if err != nil {
			http.Error(w, "failed to retrieve data from the database", http.StatusBadRequest)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
