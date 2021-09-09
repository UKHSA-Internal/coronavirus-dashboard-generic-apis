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
		"ltla":          "E09000033",
		"ltlaName":      "Westminster",
		"msoa":          "E02000979",
		"msoaName":      "Central Westminster",
		"nation":        "E92000001",
		"nationName":    "England",
		"nhsRegion":     "E40000003",
		"nhsRegionName": "London",
		"nhsTrust":      "RJ1",
		"nhsTrustName":  "Guy's and St Thomas' NHS Foundation Trust",
		"region":        "E12000007",
		"regionName":    "London",
		"utla":          "E09000033",
		"utlaName":      "Westminster",
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
