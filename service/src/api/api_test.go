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

} // TestFromDataBase

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
		{"areaCode": "E06000041", "areaName": "Wokingham", "areaType": "utla"},
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

} // TestFromDataBase
