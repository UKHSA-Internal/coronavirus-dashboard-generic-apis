package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
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

func checkResponse(t *testing.T, topic string, expected, actual interface{}) {

	if expected != actual {
		t.Errorf("[%v] Expected response code <%v>. Got <%v>\n", topic, expected, actual)
	}

} // checkResponse

func TestPostcode(t *testing.T) {

	api.Initialize()

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

	url, _ := api.Router.Get("code").URL("area_type", tokens[0].topic, "area_code", tokens[0].expected)

	req, _ := http.NewRequest("GET", url.String(), nil)
	response := executeRequest(req)

	checkResponse(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	if err := json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err.Error())
	}

	for _, itemToken := range tokens {
		checkResponse(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
	}

} // TestPostcode

func TestRegion(t *testing.T) {

	api.Initialize()

	tokens := testTokens{
		{"region", "E12000007", false},
		{"regionName", "London", true},
		{"nation", "E92000001", false},
		{"nationName", "England", true},
	}

	url, err := api.Router.Get("code").URL("area_type", tokens[0].topic, "area_code", tokens[0].expected)
	if err != nil {
		panic(err.Error())
	}
	req, _ := http.NewRequest("GET", url.String(), nil)
	response := executeRequest(req)

	checkResponse(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	err = json.Unmarshal(response.Body.Bytes(), &data)
	if err != nil {
		t.Error(err.Error())
	}

	for _, itemToken := range tokens {
		checkResponse(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
	}
} // TestRegion

func TestUtla(t *testing.T) {

	api.Initialize()

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

	url, _ := api.Router.Get("code").URL("area_type", tokens[0].topic, "area_code", tokens[0].expected)

	req, _ := http.NewRequest("GET", url.String(), nil)
	response := executeRequest(req)

	checkResponse(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	if err := json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err.Error())
	}

	for _, itemToken := range tokens {
		checkResponse(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
	}

} // TestUtla

func TestMsoa(t *testing.T) {

	api.Initialize()

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

	url, _ := api.Router.Get("code").URL("area_type", tokens[0].topic, "area_code", tokens[0].expected)

	req, _ := http.NewRequest("GET", url.String(), nil)
	response := executeRequest(req)

	checkResponse(t, "responseCode", http.StatusOK, response.Code)

	data := make(map[string]interface{})
	if err := json.Unmarshal(response.Body.Bytes(), &data); err != nil {
		t.Error(err.Error())
	}

	for _, itemToken := range tokens {
		checkResponse(t, itemToken.topic, itemToken.expected, data[itemToken.topic])
	}

} // TestMsoa

func TestAreaByType(t *testing.T) {

	api.Initialize()

	url, _ := api.Router.Get("area").URL("area_type", "nation")

	req, _ := http.NewRequest("GET", url.String(), nil)
	response := executeRequest(req)
	fmt.Println(response.Body)
	checkResponse(t, "responseCode", http.StatusOK, response.Code)

} // TestAreaByType
