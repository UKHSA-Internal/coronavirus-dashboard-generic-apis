package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"generic_apis/assert"
	"generic_apis/base"
	"generic_apis/db"
	"generic_apis/insight"

	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type (
	testToken struct {
		topic        string
		expected     string
		responseOnly bool
	}

	testTokens []testToken
)

var api = &base.Api{
	Routes: UrlPatterns,
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	rr := httptest.NewRecorder()
	api.Router.ServeHTTP(rr, req)

	return rr

} // executeRequest

func TestPostcode(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	tokens := testTokens{
		{"postcode", "SW1A 0AA", false},
		{"trimmedPostcode", "SW1A0AA", true},
		{"msoa", "E02000979", false},
		{"msoaName", "Central Westminster", true},
		{"utla", "E09000033", false},
		{"utlaName", "Westminster", true},
		{"ltla", "E09000033", false},
		{"ltlaName", "Westminster", true},
		{"nhsRegion", "E40000003", false},
		{"nhsRegionName", "London", true},
		{"nhsTrust", "RJ1", false},
		{"nhsTrustName", "Guy's and St Thomas' NHS Foundation Trust", true},
		{"region", "E12000007", false},
		{"regionName", "London", true},
		{"nation", "E92000001", false},
		{"nationName", "England", true},
	}

	url, err := api.Router.Get("code").URL("area_type", tokens[0].topic, "area_code", tokens[0].expected)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}

	response := executeRequest(req)

	fmt.Println(response.Body.String())

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	if err := json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err.Error())
	}

	t.Run("Test level payloads", func(t *testing.T) {
		t.Parallel()

		for _, itemToken := range tokens {
			assert.Equal(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
		}
	})

} // TestPostcode

func TestRegion(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	tokens := testTokens{
		{"region", "E12000007", false},
		{"regionName", "London", true},
		{"nation", "E92000001", false},
		{"nationName", "England", true},
	}

	url, err := api.Router.Get("code").URL("area_type", tokens[0].topic, "area_code", tokens[0].expected)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}

	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	err = json.Unmarshal(response.Body.Bytes(), &data)
	if err != nil {
		t.Error(err)
	}

	t.Run("Test level payloads", func(t *testing.T) {
		t.Parallel()

		for _, itemToken := range tokens {
			assert.Equal(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
		}
	})

} // TestRegion

func TestUtla(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	tokens := testTokens{
		{"utla", "E09000033", false},
		{"utlaName", "Westminster", true},
		{"nhsRegion", "E40000003", false},
		{"nhsRegionName", "London", true},
		{"region", "E12000007", false},
		{"regionName", "London", true},
		{"nation", "E92000001", false},
		{"nationName", "England", true},
	}

	url, err := api.Router.Get("code").URL("area_type", tokens[0].topic, "area_code", tokens[0].expected)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	if err := json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	t.Run("Test level payloads", func(t *testing.T) {
		t.Parallel()

		for _, itemToken := range tokens {
			assert.Equal(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
		}
	})

} // TestUtla

func TestAreaName(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	expected := map[string]interface{}{
		"areaCode": "E07000027", "areaType": "ltla", "areaName": "Barrow-in-Furness",
	}

	url, err := api.Router.Get("area_name").URL("area_type", "ltla", "area_name", "Barrow-in-Furness")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	assert.JsonObjResponseMatchExpected(t, expected, response.Body.Bytes())

} // TestPageAreaQuery

func TestMsoa(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	tokens := testTokens{
		{"msoa", "E02000979", false},
		{"msoaName", "Central Westminster", true},
		{"utla", "E09000033", false},
		{"utlaName", "Westminster", true},
		{"ltla", "E09000033", false},
		{"ltlaName", "Westminster", true},
		{"nhsRegion", "E40000003", false},
		{"nhsRegionName", "London", true},
		{"nhsTrust", "RJ1", false},
		{"nhsTrustName", "Guy's and St Thomas' NHS Foundation Trust", true},
		{"region", "E12000007", false},
		{"regionName", "London", true},
		{"nation", "E92000001", false},
		{"nationName", "England", true},
	}

	url, err := api.Router.Get("code").URL("area_type", tokens[0].topic, "area_code", tokens[0].expected)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	if err := json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	t.Run("Test level payloads", func(t *testing.T) {
		t.Parallel()

		for _, itemToken := range tokens {
			assert.Equal(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
		}
	})

} // TestMsoa

// func _TestAreaByType(t *testing.T) {
//
// 	var err error
// 	api.Insight = insight.InitialiseInsightClient()
// 	defer appinsights.TrackPanic(api.Insight, true)
//
// 	api.Database, err = db.Connect(api.Insight)
// 	if err != nil {
// 		panic(err)
// 	}
// 	api.Initialize()
// 	defer api.Database.CloseConnection()
//
// 	url, err := api.Router.Get("area").URL("area_type", "nation")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fmt.Println(">>>>", url)
//
// 	req, err := http.NewRequest("GET", url.String(), nil)
// 	if err != nil {
// 		t.Error(err)
// 	}
//
// 	response := executeRequest(req)
// 	fmt.Println(response.Body.String())
// 	assert.Equal(t, "responseCode", http.StatusOK, response.Code)
//
// } // TestAreaByType

func TestPageAreaQuery(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	expected := []map[string]interface{}{
		{"areaCode": "E92000001", "areaType": "nation", "areaName": "England"},
		{"areaCode": "N92000002", "areaType": "nation", "areaName": "Northern Ireland"},
		{"areaCode": "S92000003", "areaType": "nation", "areaName": "Scotland"},
		{"areaCode": "W92000004", "areaType": "nation", "areaName": "Wales"},
	}

	url, err := api.Router.Get("page_areas_with_type").URL("page", "Cases", "area_type", "nation")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	assert.JsonArrResponseMatchExpected(t, expected, response.Body.Bytes())

} // TestPageAreaQuery

func TestAreaOnlyQuery(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	expected := []map[string]interface{}{
		{"areaCode": "E92000001", "areaType": "nation", "areaName": "England"},
		{"areaCode": "N92000002", "areaType": "nation", "areaName": "Northern Ireland"},
		{"areaCode": "S92000003", "areaType": "nation", "areaName": "Scotland"},
		{"areaCode": "W92000004", "areaType": "nation", "areaName": "Wales"},
		{"areaCode": "E10000002", "areaName": "Buckinghamshire", "areaType": "utla"},
		{"areaCode": "E07000119", "areaName": "Fylde", "areaType": "ltla"},
		{"areaCode": "E12000004", "areaName": "East Midlands", "areaType": "region"},
	}

	url, err := api.Router.Get("page_areas").URL("page", "Cases")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	assert.JsonArrResponseContains(t, expected, response.Body.Bytes())

} // TestAreaOnlyQuery

func TestMetricSearchQuery(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

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

	url, err := api.Router.Get("metric_search").Queries("search", "casesBySpecimen").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	type cc map[string]interface{}
	data := make([]cc, 11)
	_ = json.Unmarshal(response.Body.Bytes(), &data)

	assert.JsonArrResponseContains(t, expected, response.Body.Bytes())

} // TestMetricSearchQuery

func TestMetricSearchQueryEmptyResponse(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	expected := "[]"

	url, err := api.Router.Get("metric_search").Queries("search", "invalidinput").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)
	assert.Equal(t, "empty response", response.Body.String(), expected)

} // TestMetricSearchQueryEmptyResponse

func TestMetricSearchQueryWithCategory(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

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

	url, err := api.Router.Get("metric_search").Queries("search", "casesBySpecimen", "category", "cases").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	type cc map[string]interface{}
	data := make([]cc, 11)
	_ = json.Unmarshal(response.Body.Bytes(), &data)

	assert.JsonArrResponseContains(t, expected, response.Body.Bytes())

} // TestMetricSearchQueryWithCategory

func TestMetricSearchQueryWithTags(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	expected := []map[string]interface{}{
		{"metric": "cumCasesBySpecimenDateRate", "metric_name": "Cumulative cases by specimen date rate", "category": "Cases", "tags": []interface{}{"cumulative", "event date", "incidence rate"}},
	}

	url, err := api.Router.Get("metric_search").Queries("search", "casesBySpecimen", "tags", "incidence rate").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	type cc map[string]interface{}
	data := make([]cc, 11)
	_ = json.Unmarshal(response.Body.Bytes(), &data)

	assert.JsonArrResponseContains(t, expected, response.Body.Bytes())

} // TestMetricSearchQueryWithTags

func TestMetricSearchQueryWithCategoryAndTag(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	expected := []map[string]interface{}{
		{"metric": "cumCasesBySpecimenDate", "metric_name": "Cumulative cases by specimen date", "category": "Cases", "tags": []interface{}{"cumulative", "event date"}},
		{"metric": "cumCasesBySpecimenDateRate", "metric_name": "Cumulative cases by specimen date rate", "category": "Cases", "tags": []interface{}{"cumulative", "event date", "incidence rate"}},
		{"metric": "newCasesBySpecimenDateRollingRate", "metric_name": "New cases by specimen date rolling rate", "category": "Cases", "tags": []interface{}{"daily", "event date", "prevalence rate"}},
	}

	url, err := api.Router.Get("metric_search").Queries("search", "casesBySpecimen", "category", "cases", "tags", "cumulative,prevalence rate").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	type cc map[string]interface{}
	data := make([]cc, 11)
	_ = json.Unmarshal(response.Body.Bytes(), &data)

	assert.JsonArrResponseContains(t, expected, response.Body.Bytes())

} // TestMetricSearchQueryWithCategoryAndTag

func TestMetricSearchQueryByTagOnly(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	expected := 9

	url, err := api.Router.Get("metric_search").Queries("tags", "prevalence rate").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	type cc map[string]interface{}
	data := make([]cc, 11)
	_ = json.Unmarshal(response.Body.Bytes(), &data)

	assert.Equal(t, "response length", len(data), expected)

} // TestMetricSearchQueryEmptyResponse

func TestMetricSearchQueryByCategoryOnly(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	expected := 14

	url, err := api.Router.Get("metric_search").Queries("category", "healthcare").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	type cc map[string]interface{}
	data := make([]cc, 11)
	_ = json.Unmarshal(response.Body.Bytes(), &data)

	assert.Equal(t, "response length", len(data), expected)

} // TestMetricSearchQueryByCategoryOnly

func TestMetricSearchQueryByCategoryAndTagsOnly(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	expected := 1

	url, err := api.Router.Get("metric_search").Queries("category", "healthcare", "tags", "prevalence rate").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	type cc map[string]interface{}
	data := make([]cc, 11)
	_ = json.Unmarshal(response.Body.Bytes(), &data)

	assert.Equal(t, "response length", len(data), expected)

} // TestMetricSearchQueryByCategoryAndTagsOnly

func TestMetricPropsByCategory(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.Get("metric_props").Queries("by", "category").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make([]map[string]interface{}, 0)
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	expectKey := []string{"Vaccinations", "Cases", "Deaths", "Healthcare", "Testing"}

	for _, item := range data {
		if _, ok := item["category"]; !ok {
			t.Errorf("response item does not contain 'category' in JSON object %v", item)
		} else if _, ok = item["payload"]; !ok {
			t.Errorf("response item does not contain 'payload' in JSON object %v", item)
		}

		itemCategory := item["category"]
		found := false

		for _, key := range expectKey {
			if key == itemCategory {
				found = true
			}
		}

		if !found {
			t.Errorf("key: '%s' not found in JSON object %v", itemCategory, data)
		}
	}

} // TestMetricPropsByCategory

func TestMetricPropsByTag(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.Get("metric_props").Queries("by", "tag").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make([]map[string]interface{}, 0)
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	expectKey := []string{
		"cumulative", "daily", "event date", "incidence rate",
		"national statistics", "prevalence rate", "reporting date", "weekly",
	}

	for _, item := range data {
		if _, ok := item["tag"]; !ok {
			t.Errorf("response item does not contain 'tag' in JSON object %v", item)
		} else if _, ok = item["payload"]; !ok {
			t.Errorf("response item does not contain 'payload' in JSON object %v", item)
		}

		itemTag := item["tag"]
		found := false

		for _, key := range expectKey {
			if key == itemTag {
				found = true
			}
		}

		if !found {
			t.Errorf("key: '%s' not found in JSON object %v", itemTag, response)
		}
	}

} // TestMetricPropsByTag

func TestChangeLog(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.Get("change_logs").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)
	data := make(map[string]interface{})
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	t.Run("Request tests", func(t *testing.T) {
		t.Parallel()

		expectedPage := 1
		assert.Equal(t, "current page", int(data["page"].(float64)), expectedPage)

		expectedLen := int(data["length"].(float64))
		assert.Equal(t, "payload count", len(data["data"].([]interface{})), expectedLen)
	})

} // TestChangeLog

func TestPaginatedChangeLog(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.Get("change_logs").Queries("page", "2").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)
	data := make(map[string]interface{})
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	t.Run("Request tests", func(t *testing.T) {
		t.Parallel()

		expectedPage := 2
		assert.Equal(t, "current page", int(data["page"].(float64)), expectedPage)

		expectedLen := len(data["data"].([]interface{}))
		assert.Equal(t, "payload count", int(data["length"].(float64)), expectedLen)
	})

} // TestPaginatedChangeLog

func TestChangeLogSearch(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.Get("change_logs").Queries("search", "incor").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)
	data := make(map[string]interface{})
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	t.Run("Request tests", func(t *testing.T) {
		t.Parallel()

		expectedPage := 1
		assert.Equal(t, "current page", int(data["page"].(float64)), expectedPage)

		returnedLen := len(data["data"].([]interface{}))
		assert.Equal(t, "payload count", int(data["length"].(float64)), returnedLen)

		expectedLenGreaterThan := 2
		assert.IntGreater(t, "payload count", int(data["length"].(float64)), expectedLenGreaterThan)
	})

} // TestChangeLogSearch

func TestDatedChangeLog(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.Get("change_logs_single_month").URL("date", "2021-05")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	expected := 17 // Expected length for May 2021.
	assert.Equal(t, "response length", int(data["length"].(float64)), expected)

} // TestDatedChangeLog

func TestDatedChangeLogSearch(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	monthParam := "2021-05"
	monthParamLen := len(monthParam)

	url, err := api.Router.
		Get("change_logs_single_month").
		Queries("search", "cases in england").
		URL("date", monthParam)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	expected := 5 // Expected length for May 2021 with search query "cases in England".
	assert.Equal(t, "response length", int(data["length"].(float64)), expected)

	// Asserting that the data of each record is consistent
	// with that which is requested.
	for _, record := range data["data"].([]interface{}) {
		assert.Equal(
			t,
			"date match",
			monthParam,
			record.(map[string]interface{})["date"].(string)[:monthParamLen],
		)
	}

} // TestDatedChangeLogSearch

func TestLogBanner(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.
		Get("log_banners").
		URL("date", "2021-05-21", "page", "Cases", "area_type", "overview", "area_name", "United Kingdom")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make([]interface{}, 0)
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	expected := 1 // Expected length for May 2021 with search query "cases in England".
	assert.Equal(t, "response length", len(data), expected)

} // TestLogBanner

func TestChangeLogDates(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.Get("change_logs_components").URL("component", "dates")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make([]interface{}, 0)
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	expected := 1
	assert.IntGreater(t, "response length", len(data), expected)

} // TestChangeLogDates

func TestChangeLogTypes(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.Get("change_logs_components").URL("component", "types")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make([]interface{}, 0)
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	expected := 1
	assert.IntGreater(t, "response length", len(data), expected)

} // TestChangeLogTypes

func TestChangeLogTitles(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.Get("change_logs_components").URL("component", "titles")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make([]interface{}, 0)
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	expected := 1
	assert.IntGreater(t, "response length", len(data), expected)

} // TestChangeLogTitles

func TestChangeLogItem(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.
		Get("change_log_item").
		URL("id", "3e80c9b6-ebc5-4669-9e94-80f894e6fd38")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	data := map[string]interface{}{}
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	expected := "3e80c9b6-ebc5-4669-9e94-80f894e6fd38"

	assert.Equal(t, "IDs match", expected, data["id"].(string))

} // TestChangeLogItem

func TestAnnouncements(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.Get("announcements").URL()
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make([]interface{}, 0)
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	expected := 10
	assert.IntGreater(t, "response length", len(data), expected)

} // TestAnnouncements

func TestAnnouncementItem(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.
		Get("announcement_item").
		URL("id", "e6817475-dd68-4153-859c-91bc209421e6")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	expected := map[string]interface{}{
		"body": "On Thursday 5 August, case, hospital admission and death rates per 100,000 people for the UK, " +
			"nations, regions and local authorities will be updated to use the mid-2020 population estimates.",
		"date":        "2021-08-03T00:00:00Z",
		"expire":      "2021-08-05",
		"guid":        "e6817475-dd68-4153-859c-91bc209421e6",
		"has_expired": true,
		"launch":      "2021-08-03",
	}

	assert.JsonObjResponseMatchExpected(t, expected, response.Body.Bytes())

} // TestAnnouncementItem

func TestMetricDoc(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	url, err := api.Router.
		Get("metric_doc").
		URL("metric", "newCasesBySpecimenDate")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{}, 0)
	if err = json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err)
	}

	assert.Equal(t, "metrics match", "newCasesBySpecimenDate", data["metric"].(string))

} // TestMetricDoc

func TestMetricAreas(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.Database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.Database.CloseConnection()

	expected := []map[string]interface{}{
		{"area_type": "ltla", "last_update": "2021-08-25"},
		{"area_type": "msoa", "last_update": "2021-08-21"},
		{"area_type": "nation", "last_update": "2021-08-25"},
		{"area_type": "overview", "last_update": "2021-08-25"},
		{"area_type": "region", "last_update": "2021-08-25"},
		{"area_type": "utla", "last_update": "2021-08-25"},
	}

	url, err := api.Router.
		Get("metric_areas").
		URL("metric", "newCasesBySpecimenDate", "date", "2021-08-26")
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		t.Error(err)
	}
	response := executeRequest(req)

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	assert.JsonArrResponseMatchExpected(t, expected, response.Body.Bytes())

} // TestMetricAreas
