package area_name

import (
	"testing"

	"generic_apis/apps/utils"
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

	expected := map[string]interface{}{
		"areaCode": "E02000979",
		"areaName": "Central Westminster",
		"areaType": "msoa",
	}

	response, err := conf.fromDatabase("msoa", "Central Westminster")
	if err != nil {
		t.Error(err.Error())
	}

	jsonResponse, err := utils.JSONMarshal(response)
	if err != nil {
		t.Error(err.Error())
	}

	assert.JsonObjResponseMatchExpected(t, expected, jsonResponse)

} // TestFromDataBase
