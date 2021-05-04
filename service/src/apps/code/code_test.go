package code

import (
	"testing"

	"generic_apis/db"
	"generic_apis/insight"
	"generic_apis/testify"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func TestFromDataBase(t *testing.T) {

	insightClient := insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(insightClient, true)

	database, err := db.Connect(insightClient)
	if err != nil {
		panic(err)
	}
	defer database.CloseConnection()
	conf := &handler{database, ""}

	expected := map[string]interface{}{
		"ltla":            "E09000033",
		"ltlaName":        "Westminster",
		"msoa":            "E02000979",
		"msoaName":        "Central Westminster",
		"nation":          "E92000001",
		"nationName":      "England",
		"nhsRegion":       "E40000003",
		"nhsRegionName":   "London",
		"nhsTrust":        "RJ1",
		"nhsTrustName":    "Guy's and St Thomas' NHS Foundation Trust",
		"postcode":        "SW1A 0AA",
		"region":          "E12000007",
		"regionName":      "London",
		"trimmedPostcode": "SW1A0AA",
		"utla":            "E09000033",
		"utlaName":        "Westminster",
	}

	jsonResponse, err := conf.fromDatabase("postcode", "SW1A 0AA")
	if err != nil {
		// fmt.Println(err.Error())
		t.Error(err.Error())
	}

	testify.AssertJsonObjResponseMatchExpected(t, expected, jsonResponse)

} // TestFromDataBase
