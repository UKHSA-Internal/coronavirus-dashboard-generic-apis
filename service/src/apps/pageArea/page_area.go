package pageArea

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

func (conf *handler) fromDatabase(params map[string]string) ([]byte, error) {

	var (
		ok           bool
		areaType     string
		args         = []interface{}{strings.ToLower(params["page"])}
		preppedQuery = query
	)

	if areaType, ok = params["area_type"]; ok {
		areaTypeTemplate := "%s"
		if areaType != "la" {
			areaTypeTemplate = "'%s'"
		}

		areaType = strings.ToLower(areaType)

		if areaType, ok = utils.AreaTypes[areaType]; !ok {
			return nil, fmt.Errorf("invalid area type")
		} else {
			areaType = fmt.Sprintf(areaTypeTemplate, areaType)
			preppedQuery += fmt.Sprintf(areaTypeFilter, areaType)
		}
	}

	preppedQuery += queryExtras

	payload := &db.Payload{
		Query:         fmt.Sprintf(queryWrapper, preppedQuery),
		Args:          args,
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
		defer conf.db.CloseConnection()

		pathVars := mux.Vars(r)

		response, err := conf.fromDatabase(pathVars)
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(response); err != nil {
			panic(err)
		}

	}

} // queryByCode
