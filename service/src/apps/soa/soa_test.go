package soa

import (
	"encoding/json"
	"reflect"
	"testing"
)

type soaResponse struct {
	AreaCode         string  `json:"areaCode"`
	AreaName         string  `json:"areaName"`
	AreaType         string  `json:"areaType"`
	Change           int16   `json:"change"`
	ChangePercentage int16   `json:"changePercentage"`
	Date             string  `json:"date"`
	Direction        string  `json:"direction"`
	RollingRate      float32 `json:"rollingRate"`
	RollingSum       int16   `json:"rollingSum"`
} // soaResponse

func TestTimestamp(t *testing.T) {

	timestamp, err := getLatestTimestamp("region")
	if err != nil {
		t.Error(err.Error())
	}

	if reflect.TypeOf(timestamp).Kind() != reflect.String {
		t.Error("type mismatch")
	}

} // TestTimestamp

func checkResponse(t *testing.T, expected, actual interface{}) {

	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}

} // checkResponse

func TestFromDatabase(t *testing.T) {

	result, err := fromDatabase("msoa", "E02000408")
	if err != nil {
		t.Error(err.Error())
	}

	var jsonData soaResponse
	if err = json.Unmarshal(result, &jsonData); err != nil {
		t.Error(err.Error())
	}

	checkResponse(t, "E02000408", jsonData.AreaCode)
	checkResponse(t, "Tottenham North West", jsonData.AreaName)
	checkResponse(t, "msoa", jsonData.AreaType)

} // TestFromDatabase
