package pageArea

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

const (
	query = `
SELECT area_type AS "areaType", area_name AS "areaName", area_code AS "areaCode"
FROM covid19.page AS pg
	JOIN covid19.page_area_reference AS par ON par.category_id = pg.id
	JOIN covid19.area_reference AS ar ON ar.id = par.area_id
WHERE LOWER(pg.title) = $1`

	areaTypeFilter = ` AND ar.area_type = $2`
)

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

func (conf *handler) fromDatabase(params map[string]string) ([]byte, error) {

	var (
		ok           bool
		areaType     string
		args         = []interface{}{strings.ToLower(params["page"])}
		preppedQuery = query
	)

	if areaType, ok = params["area_type"]; ok {
		areaType = strings.ToLower(areaType)

		if areaType, ok = areaTypes[areaType]; !ok {
			return nil, fmt.Errorf("invalid area type")
		} else {
			args = append(args, areaType)
			preppedQuery += areaTypeFilter
		}
	}

	payload := &db.Payload{
		Query:         preppedQuery,
		Args:          args,
		OperationData: insight.GetOperationData(conf.traceparent),
	}
	results, err := conf.db.FetchAll(payload)
	if err != nil {
		return nil, err
	}

	if jsonString, err := json.Marshal(results); err != nil {
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

		if _, ok := pathVars["page"]; !ok {
			panic("no page")
		}

		response, err := conf.fromDatabase(pathVars)
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
