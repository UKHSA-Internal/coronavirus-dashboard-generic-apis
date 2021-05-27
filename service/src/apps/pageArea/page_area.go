package pageArea

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

func (conf *handler) fromDatabase(params map[string]string) ([]byte, error) {

	var (
		ok           bool
		areaType     string
		args         = []interface{}{strings.ToLower(params["page"])}
		preppedQuery = query
	)

	if areaType, ok = params["area_type"]; ok {
		areaType = strings.ToLower(areaType)

		if areaType, ok = utils.AreaTypes[areaType]; !ok {
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

	return json.Marshal(results)

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
