package area_by_type

import (
	"testing"

	"generic_apis/assert"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func TestFromDataBase(t *testing.T) {

	t.Parallel()

	insightClient := insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(insightClient, true)

	database, err := db.Connect(insightClient)
	if err != nil {
		panic(err)
	}
	defer database.CloseConnection()
	conf := &handler{database, ""}

	expected := []map[string]interface{}{
		{"areaCode": "E92000001", "areaName": "England"},
		{"areaCode": "N92000002", "areaName": "Northern Ireland"},
		{"areaCode": "S92000003", "areaName": "Scotland"},
		{"areaCode": "W92000004", "areaName": "Wales"},
	}

	jsonResponse, err := conf.fromDatabase("nation")
	if err != nil {
		t.Error(err)
	}

	assert.JsonArrResponseMatchExpected(t, expected, jsonResponse)

} // TestFromDataBase
