package change_logs

import (
	"net/http"

	"generic_apis/apps/utils"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/pkg/errors"
)

func (conf *handler) getComponentsFromDatabase(component string) ([]db.ResultType, *utils.FailedResponse) {

	failure := &utils.FailedResponse{}

	payload := &db.Payload{
		Query:         componentQueries[component],
		Args:          []interface{}{},
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	res, err := conf.db.FetchAll(payload)
	if err != nil {
		failure.Response = errors.Errorf("failed to retrieve data")
		failure.HttpCode = http.StatusInternalServerError
		failure.Payload = err

		return nil, failure
	}
	return res, nil

} // getDatesFromDatabase
