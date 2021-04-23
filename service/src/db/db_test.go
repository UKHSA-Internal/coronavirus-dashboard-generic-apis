package db

import (
	"testing"
)

func TestDatabase(t *testing.T) {

	query := `SELECT 1 AS value`

	results, err := Query(query)
	if err != nil {
		t.Error(err)
	}

	if len(results) < 1 {
		t.Errorf("invalid response - length smaller than one")
	}

	if results[0]["value"].(int32) != 1 {
		t.Errorf("invalid value")
	}

} // TestDatabase
