package areaByType

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"generic_apis/db"
	"github.com/gorilla/mux"
)

const query = `
SELECT area_name, area_code
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

func fromDatabase(areaType string) ([]byte, error) {

	var ok bool
	if areaType, ok = areaTypes[strings.ToLower(areaType)]; !ok {
		return nil, fmt.Errorf("invalid area type")
	}

	results, err := db.Query(query, areaType)
	if err != nil {
		return nil, err
	}

	if jsonString, err := json.Marshal(results); err != nil {
		return nil, err
	} else {
		return jsonString, nil
	}

} // FromDatabase

func AreaByTypeQuery(w http.ResponseWriter, r *http.Request) {

	pathVars := mux.Vars(r)

	var (
		areaType string
		ok       bool
	)

	if areaType, ok = pathVars["area_type"]; !ok {
		panic("no area type")
	}

	response, err := fromDatabase(areaType)
	if err != nil {
		panic(err)
	}

	if _, err := w.Write(response); err != nil {
		panic(err)
	}

} // queryByCode
