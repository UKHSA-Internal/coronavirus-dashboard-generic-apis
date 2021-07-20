package change_logs

import (
	"generic_apis/db"
	"generic_apis/insight"
)

func (conf *handler) getComponentsFromDatabase(component string) ([]db.ResultType, error) {

	payload := &db.Payload{
		Query:         componentQueries[component],
		Args:          []interface{}{},
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	res, err := conf.db.FetchAll(payload)

	return res, err

} // getDatesFromDatabase
