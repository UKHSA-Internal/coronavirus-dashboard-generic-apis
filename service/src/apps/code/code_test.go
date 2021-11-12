package code

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
		"ltla":              "E09000033",
		"ltlaCentroid":      []interface{}{-0.157209256, 51.514010951},
		"ltlaName":          "Westminster",
		"msoa":              "E02000979",
		"msoaCentroid":      []interface{}{-0.133169461, 51.49905812},
		"msoaName":          "Central Westminster",
		"nation":            "E92000001",
		"nationCentroid":    []interface{}{-1.463684694, 52.592370945},
		"nationName":        "England",
		"nhsRegion":         "E40000003",
		"nhsRegionCentroid": []interface{}{},
		"nhsRegionName":     "London",
		"nhsTrust":          "RJ1",
		"nhsTrustCentroid":  []interface{}{},
		"nhsTrustName":      "Guy's and St Thomas' NHS Foundation Trust",
		"postcode":          "SW1A 0AA",
		"region":            "E12000007",
		"regionCentroid":    []interface{}{-0.110578143, 51.500891693},
		"regionName":        "London",
		"trimmedPostcode":   "SW1A0AA",
		"utla":              "E09000033",
		"utlaCentroid":      []interface{}{-0.160411474, 51.513701283},
		"utlaName":          "Westminster",
	}

	response, err := conf.fromDatabase("postcode", "SW1A 0AA")
	if err != nil {
		t.Error(err.Error())
	}
	jsonResponse, err := utils.JSONMarshal(response)
	if err != nil {
		t.Error(err.Error())
	}

	assert.JsonObjResponseMatchExpected(t, expected, jsonResponse)

} // TestFromDataBase
