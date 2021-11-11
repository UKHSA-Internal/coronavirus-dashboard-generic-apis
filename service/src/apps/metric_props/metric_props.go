package metric_props

import (
	"net/http"
	"net/url"
	"strings"

	"generic_apis/apps/utils"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type handler struct {
	db          *db.Config
	insight     appinsights.TelemetryClient
	traceparent string
}

func (conf *handler) fromDatabase(params url.Values) ([]db.ResultType, error) {

	var (
		preppedQuery string
		args         []interface{}
	)

	switch params.Get("by") {
	case "category":
		preppedQuery = byCategory
		break
	case "tag":
		preppedQuery = byTag
		break
	case "areaType":
		req := &utils.GenericRequest{
			Traceparent: conf.traceparent,
			Insight:     conf.insight,
		}
		if timestamp, err := req.GetLatestTimeStamp(); err != nil {
			panic(err)
		} else {
			preppedQuery = strings.ReplaceAll(byAreaType, partitionDatePlaceholder, timestamp)
		}
		break
	default:
		panic("invalid value for 'by'")
	}

	payload := &db.Payload{
		Query:         preppedQuery,
		Args:          args,
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	return conf.db.FetchAll(payload)

} // fromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{insight: insight}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			err         error
			response    []db.ResultType
			jsonPayload []byte
		)

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(insight)
		if err != nil {
			panic(err)
		}

		response, err = conf.fromDatabase(r.URL.Query())
		if err != nil {
			panic(err)
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
