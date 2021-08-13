package page_area

import (
	"testing"

	"generic_apis/apps/utils"
	"generic_apis/assert"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func TestPageAreaQuery(t *testing.T) {

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
		{"areaCode": "E92000001", "areaType": "nation", "areaName": "England"},
		{"areaCode": "N92000002", "areaType": "nation", "areaName": "Northern Ireland"},
		{"areaCode": "S92000003", "areaType": "nation", "areaName": "Scotland"},
		{"areaCode": "W92000004", "areaType": "nation", "areaName": "Wales"},
	}

	params := map[string]string{
		"page":      "Cases",
		"area_type": "nation",
	}

	response, err := conf.fromDatabase(params)
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err.Error())
	}
	jsonResponse, err := utils.JSONMarshal(response)
	if err != nil {
		t.Error(err.Error())
	}

	assert.JsonArrResponseMatchExpected(t, expected, jsonResponse)

} // TestFromDataBase

func TestAreaOnlyQuery(t *testing.T) {

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
		{"areaCode": "E92000001", "areaType": "nation", "areaName": "England"},
		{"areaCode": "N92000002", "areaType": "nation", "areaName": "Northern Ireland"},
		{"areaCode": "S92000003", "areaType": "nation", "areaName": "Scotland"},
		{"areaCode": "W92000004", "areaType": "nation", "areaName": "Wales"},
		{"areaCode": "E10000002", "areaName": "Buckinghamshire", "areaType": "utla"},
		{"areaCode": "E07000119", "areaName": "Fylde", "areaType": "ltla"},
		{"areaCode": "E12000004", "areaName": "East Midlands", "areaType": "region"},
	}

	params := map[string]string{
		"page": "Deaths",
	}

	response, err := conf.fromDatabase(params)
	if err != nil {
		t.Error(err)
	}
	if err != nil {
		t.Error(err.Error())
	}
	jsonResponse, err := utils.JSONMarshal(response)
	if err != nil {
		t.Error(err.Error())
	}

	assert.JsonArrResponseContains(t, expected, jsonResponse)

} // TestFromDataBase
