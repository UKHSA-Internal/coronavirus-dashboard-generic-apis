package assert

import (
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func JsonObjResponseMatchExpected(t *testing.T, expected interface{}, jsonResponse []byte) {

	t.Run("JsonObjResponseMatchExpected", func(t *testing.T) {
		response := make(map[string]interface{})
		if err := json.Unmarshal(jsonResponse, &response); err != nil {
			t.Error(err)
		}

		if !cmp.Equal(response, expected) {
			difference := cmp.Diff(response, expected)
			t.Errorf(
				"response does not match the expected output\nExpected: %v\nActual: %v\nDifference: %v",
				expected, response, difference)
		}
	})

} // JsonResponseMatchExpected

func JsonArrResponseMatchExpected(t *testing.T, expected interface{}, jsonResponse []byte) {

	t.Run("JsonArrResponseMatchExpected", func(t *testing.T) {
		response := make([]map[string]interface{}, 0)
		if err := json.Unmarshal(jsonResponse, &response); err != nil {
			t.Error(err)
		}

		if !cmp.Equal(response, expected) {
			difference := cmp.Diff(response, expected)
			t.Errorf(
				"response does not match the expected output\nExpected: %v\nActual: %v\nDifference: %v",
				expected, response, difference)
		}
	})

} // JsonResponseMatchExpected

func JsonArrResponseContains(t *testing.T, expected interface{}, jsonResponse []byte) {

	t.Run("JsonArrResponseContains", func(t *testing.T) {

		response := make([]map[string]interface{}, 0)
		if err := json.Unmarshal(jsonResponse, &response); err != nil {
			t.Error(err)
		}

		lenExpected := len(expected.([]map[string]interface{}))
		lenDetected := 0

		for _, expectedItem := range expected.([]map[string]interface{}) {
			for _, responseItem := range response {
				if cmp.Equal(expectedItem, responseItem) {
					lenDetected += 1
					break
				}
			}
		}

		if lenExpected != lenDetected {
			t.Errorf("Expected to find %d items in the response, found %d instead", lenExpected, lenDetected)
		}
	})

} // JsonResponseMatchExpected

func Equal(t *testing.T, topic string, expected, actual interface{}) {

	t.Run(topic, func(t *testing.T) {
		if expected != actual {
			t.Errorf("[%v] Expected response code <%v>. Got <%v>\n", topic, expected, actual)
		}
	})

} // Equal

func IntGreater(t *testing.T, topic string, expected, actual int) {

	t.Run(topic, func(t *testing.T) {
		if expected <= actual {
			t.Errorf("[%v] Expected response code <%v>. Got <%v>\n", topic, expected, actual)
		}
	})

} // IntGreater
