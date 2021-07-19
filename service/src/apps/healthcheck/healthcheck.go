package healthcheck

import (
	"net/http"
	"time"

	hc "github.com/etherlabsio/healthcheck/v2"
	"github.com/etherlabsio/healthcheck/v2/checkers"
)

// type handler struct {
// 	db          *db.Config
// 	traceparent string
// }
//
// func (conf *handler) fromDatabase() ([]byte, error) {
//
// 	payload := &db.Payload{
// 		Query:         healthCheckQuery,
// 		Args:          []interface{}{},
// 		OperationData: insight.GetOperationData(conf.traceparent),
// 	}
//
// 	_, err := conf.db.FetchAll(payload)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	data := map[string]string{
// 		"database": "healthy",
// 	}
//
// 	return utils.JSONMarshal(data)
//
// } // FromDatabase

func Handler() http.Handler {

	return hc.Handler(
		hc.WithTimeout(5*time.Second),
		hc.WithObserver("diskspace", checkers.DiskSpace("/var/log", 90)),
	)

} // queryByCode
