package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"generic_apis/assert"
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

var api = &Api{}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {

	rr := httptest.NewRecorder()
	api.Router.ServeHTTP(rr, req)

	return rr

} // executeRequest

func TestPostcode(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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

	assert.Equal(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	if err := json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		fmt.Println(err.Error())
		t.Error(err.Error())
	}

	for _, itemToken := range tokens {
		assert.Equal(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
	}

} // TestPostcode

func TestRegion(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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

	for _, itemToken := range tokens {
		assert.Equal(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
	}
} // TestRegion

func TestUtla(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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

	for _, itemToken := range tokens {
		assert.Equal(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
	}

} // TestUtla

func TestMsoa(t *testing.T) {

	var err error
	api.Insight = insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(api.Insight, true)

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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

	for _, itemToken := range tokens {
		assert.Equal(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
	}

} // TestMsoa

// func _TestAreaByType(t *testing.T) {
//
// 	var err error
// 	api.Insight = insight.InitialiseInsightClient()
// 	defer appinsights.TrackPanic(api.Insight, true)
//
// 	api.database, err = db.Connect(api.Insight)
// 	if err != nil {
// 		panic(err)
// 	}
// 	api.Initialize()
// 	defer api.database.CloseConnection()
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

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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

	api.database, err = db.Connect(api.Insight)
	if err != nil {
		panic(err)
	}
	api.Initialize()
	defer api.database.CloseConnection()

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
