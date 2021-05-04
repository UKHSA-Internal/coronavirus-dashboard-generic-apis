package db

import (
	"testing"

	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func TestDatabaseConnection(t *testing.T) {

	insightClient := insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(insightClient, true)

	conf, err := Connect(insightClient)
	if err != nil {
		t.Error(err)
	}

	defer conf.CloseConnection()

} // TestDatabase

// func TestFetchAll(t *testing.T) {
//
// 	database, err := Connect()
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer database.CloseConnection()
//
// 	query := `SELECT 1 AS value`
//
// 	results, err := database.FetchAll(query)
// 	if err != nil {
// 		t.Error(err.Error())
// 	}
//
// 	if len(results) < 1 {
// 		t.Errorf("invalid response - length smaller than one")
// 	}
//
// 	if results[0]["value"].(int32) != 1 {
// 		t.Errorf("invalid value")
// 	}
//
// } // TestFetchAll

func TestFetchRow(t *testing.T) {

	insightClient := insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(insightClient, true)

	database, err := Connect(insightClient)
	if err != nil {
		panic(err)
	}
	defer database.CloseConnection()

	payload := &Payload{
		Query:         `SELECT 1 AS value`,
		Args:          []interface{}{},
		OperationData: insight.GetOperationData(""),
	}

	results, err := database.FetchRow(payload)
	if err != nil {
		t.Error(err.Error())
	}

	if results["value"].(int32) != 1 {
		t.Error("invalid value")
	}

} // TestFetchRow
