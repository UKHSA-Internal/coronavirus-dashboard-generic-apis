package areaByType

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

const query = `
SELECT area_name AS "areaName", area_code AS "areaCode"
FROM covid19.area_reference
WHERE area_type = $1
ORDER BY area_name ASC;
`

var areaTypes = map[string]string{
	"postcode":  "postcode",
	"msoa":      "msoa",
	"nhstrust":  "nhsTrust",
	"nhsregion": "nhsRegion",
	"utla":      "utla",
	"ltla":      "ltla",
	"region":    "region",
	"nation":    "nation",
	"overview":  "overview",
}

func (conf *handler) fromDatabase(areaType string) ([]byte, error) {

	var ok bool
	if areaType, ok = areaTypes[strings.ToLower(areaType)]; !ok {
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

	return json.Marshal(results)

} // FromDatabase

func Handler(config *db.Config) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{config, ""}

	return func(w http.ResponseWriter, r *http.Request) {

		conf.traceparent = r.Header.Get("traceparent")

		pathVars := mux.Vars(r)

		var (
			areaType string
			ok       bool
		)

		if areaType, ok = pathVars["area_type"]; !ok {
			panic("no area type")
		}

		response, err := conf.fromDatabase(areaType)
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
