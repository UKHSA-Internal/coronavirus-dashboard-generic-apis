package code

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"generic_apis/db"
	"generic_apis/insight"
	"github.com/gorilla/mux"
)

type handler struct {
	db          *db.Config
	traceparent string
}

const areaQuery = `
SELECT
	area_code    AS "areaCode",
	area_name    AS "areaName",
	ar.area_type AS "areaType"
FROM covid19.area_reference AS ar
WHERE id IN (
	  	 SELECT parent_id
	  	 FROM covid19.area_reference AS ar2
		   JOIN covid19.area_relation_se AS pl2 ON pl2.child_id = ar2.id
	  	 WHERE area_type = $1
	  	   AND area_code = $2
	  )
   OR ( area_type = $1 AND area_code = $2 )
`

const postcodeQuery = `
SELECT postcode, 
	   area_code AS "areaCode", 
	   area_name AS "areaName", 
	   ar.area_type AS "areaType"
FROM covid19.area_reference AS ar
  JOIN covid19.postcode_lookup AS pl ON pl.area_id = ar.id
  JOIN covid19.area_priorities AS ap ON ap.area_type = ar.area_type
WHERE UPPER(REPLACE(postcode, ' ', '')) = $2
  AND priority >= (
	SELECT priority 
	FROM covid19.area_priorities
	WHERE area_type = $1
	LIMIT 1
  )
`

var areaTypes = map[string]string{
	"postcode":        "postcode",
	"trimmedpostcode": "postcode",
	"msoa":            "msoa",
	"nhstrust":        "nhsTrust",
	"nhsregion":       "nhsRegion",
	"utla":            "utla",
	"ltla":            "ltla",
	"region":          "region",
	"nation":          "nation",
	"overview":        "overview",
}

func (conf *handler) fromDatabase(areaType, search string) ([]byte, error) {

	var (
		ok     bool
		params []interface{}
		query  string
	)

	search = strings.ReplaceAll(strings.ToUpper(search), " ", "")
	areaType = strings.ToLower(areaType)

	if areaType, ok = areaTypes[areaType]; !ok {
		return nil, fmt.Errorf("invalid area type")
	}

	if areaType == areaTypes["postcode"] {
		params = []interface{}{areaTypes["msoa"], search}
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

	if len(results) > 0 && areaType == areaTypes["postcode"] {
		postcode := results[0]["postcode"].(string)
		data["postcode"] = postcode
		data["trimmedPostcode"] = strings.ReplaceAll(postcode, " ", "")
	}

	if jsonString, err := json.Marshal(data); err != nil {
		return nil, err
	} else {
		return jsonString, nil
	}

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
