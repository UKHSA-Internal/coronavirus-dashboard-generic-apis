package soa

import (
	"reflect"
	"testing"

	"generic_apis/db"
	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func TestTimestamp(t *testing.T) {

	insightClient := insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(insightClient, true)

	database, err := db.Connect(insightClient)
	if err != nil {
		panic(err)
	}
	defer database.CloseConnection()
	conf := &handler{database, ""}

	timestamp, err := conf.getLatestTimestamp("region")
	if err != nil {
		t.Error(err.Error())
	}

	if reflect.TypeOf(timestamp).Kind() != reflect.String {
		t.Error("type mismatch")
	}

} // TestTimestamp

// func checkResponse(t *testing.T, expected, actual interface{}) {
//
// 	if expected != actual {
// 		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
// 	}
//
// } // checkResponse
