package change_logs

import (
	"generic_apis/db"
	"generic_apis/insight"
)

func (conf *handler) getTypesFromDatabase() ([]db.ResultType, error) {

	payload := &db.Payload{
		Query:         recordMonths,
		Args:          []interface{}{},
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	res, err := conf.db.FetchAll(payload)

	return res, err

} // getDatesFromDatabase
