package metric_props

import (
	"net/url"
	"testing"

	"generic_apis/db"
	"generic_apis/insight"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

func TestFromDataBaseSearchByTag(t *testing.T) {

	t.Parallel()

	insightClient := insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(insightClient, true)

	database, err := db.Connect(insightClient)
	if err != nil {
		panic(err)
	}
	defer database.CloseConnection()
	conf := &handler{database, ""}

	response, err := conf.fromDatabase(url.Values{"by": []string{"tag"}})
	if err != nil {
		t.Error(err.Error())
	}

	expectKey := []string{
		"cumulative", "daily", "event date", "incidence rate",
		"national statistics", "prevalence rate", "reporting date", "weekly",
	}

	for _, item := range response {
		if _, ok := item["tag"]; !ok {
			t.Errorf("response item does not contain 'tag' in JSON object %v", response)
		} else if _, ok = item["payload"]; !ok {
			t.Errorf("response item does not contain 'payload' in JSON object %v", response)
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

} // TestFromDataBaseSearchByTag

func TestFromDataBaseSearchByCategory(t *testing.T) {

	t.Parallel()

	insightClient := insight.InitialiseInsightClient()
	defer appinsights.TrackPanic(insightClient, true)

	database, err := db.Connect(insightClient)
	if err != nil {
		panic(err)
	}
	defer database.CloseConnection()
	conf := &handler{database, ""}

	response, err := conf.fromDatabase(url.Values{"by": []string{"category"}})
	if err != nil {
		t.Error(err.Error())
	}

	expectKey := []string{"Vaccinations", "Cases", "Deaths", "Healthcare", "Testing"}

	for _, item := range response {
		if _, ok := item["category"]; !ok {
			t.Errorf("response item does not contain 'category' in JSON object %v", response)
		} else if _, ok = item["payload"]; !ok {
			t.Errorf("response item does not contain 'payload' in JSON object %v", response)
		}

		itemCategory := item["category"]
		found := false

		for _, key := range expectKey {
			if key == itemCategory {
				found = true
			}
		}

		if !found {
			t.Errorf("key: '%s' not found in JSON object %v", itemCategory, response)
		}
	}

} // TestFromDataBaseSearchByCategory
