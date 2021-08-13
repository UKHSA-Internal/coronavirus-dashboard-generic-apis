package metric_search

import (
	"net/url"
	"testing"

	"generic_apis/assert"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func TestFromDataBaseSearch(t *testing.T) {

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
		{"metric": "changeInNewCasesBySpecimenDate", "metric_name": "Change in new cases by specimen date", "category": "Cases", "tags": []interface{}{"event date"}},
		{"metric": "cumCasesBySpecimenDate", "metric_name": "Cumulative cases by specimen date", "category": "Cases", "tags": []interface{}{"cumulative", "event date"}},
		{"metric": "cumCasesBySpecimenDateRate", "metric_name": "Cumulative cases by specimen date rate", "category": "Cases", "tags": []interface{}{"cumulative", "event date", "incidence rate"}},
		{"metric": "newCasesBySpecimenDate", "metric_name": "New cases by specimen date", "category": "Cases", "tags": []interface{}{"daily", "event date"}},
		{"metric": "newCasesBySpecimenDateAgeDemographics", "metric_name": "New cases by specimen date age demographics", "category": "Cases", "tags": []interface{}{"daily", "event date"}},
		{"metric": "newCasesBySpecimenDateChange", "metric_name": "New cases by specimen date change", "category": "Cases", "tags": []interface{}{"daily", "event date"}},
		{"metric": "newCasesBySpecimenDateChangePercentage", "metric_name": "New cases by specimen date change percentage", "category": "Cases", "tags": []interface{}{"daily", "event date"}},
		{"metric": "newCasesBySpecimenDateDirection", "metric_name": "New cases by specimen date direction", "category": "Cases", "tags": []interface{}{"daily", "event date"}},
		{"metric": "newCasesBySpecimenDateRollingRate", "metric_name": "New cases by specimen date rolling rate", "category": "Cases", "tags": []interface{}{"daily", "event date", "prevalence rate"}},
		{"metric": "newCasesBySpecimenDateRollingSum", "metric_name": "New cases by specimen date rolling sum", "category": "Cases", "tags": []interface{}{"daily", "event date"}},
		{"metric": "previouslyReportedNewCasesBySpecimenDate", "metric_name": "Previously reported new cases by specimen date", "category": "Cases", "tags": []interface{}{"event date"}},
	}

	jsonResponse, err := conf.fromDatabase(url.Values{"search": []string{"casesBySpecimenDate"}})
	if err != nil {
		t.Error(err.Error())
	}

	assert.JsonArrResponseMatchExpected(t, expected, jsonResponse)

} // TestFromDataBaseSearch
