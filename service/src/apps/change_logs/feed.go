package change_logs

import (
	"fmt"
	"net/http"
	"time"

	"generic_apis/db"
	"generic_apis/feed"
	"generic_apis/feed/atom"
	"generic_apis/feed/rss"
	"generic_apis/insight"
	"github.com/gorilla/mux"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type dbResponse struct {
	payload   *[]feed.Payload
	timestamp time.Time
}

const feedCategory = "Service logs"
const apiEndpoint = "https://api.coronavirus.data.gov.uk/generic/change_logs/"
const websiteEndpoint = "https://coronavirus.data.gov.uk/details/whatsnew/"

func (conf *handler) fromDatabaseFeed() (*dbResponse, error) {

	var (
		params []interface{}
		query  = feedQuery
	)

	payload := &db.Payload{
		Query:         query,
		Args:          params,
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	res, err := conf.db.FetchAll(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data")
	}

	var data = make([]feed.Payload, len(res))

	for index, item := range res {
		data[index].Date = item["date"].(time.Time)
		data[index].Guid = &feed.Guid{Guid: item["guid"].(string), IsPermaLink: "false"}
		data[index].Title = item["title"].(string)
		data[index].Category = item["category"].(string)
		data[index].Link = "https://coronavirus.data.gov.uk/details/whatsnew/record/" + item["guid"].(string)
		data[index].Description = item["description"].(string)
	}

	response := &dbResponse{payload: &data}

	if len(res) > 0 {
		response.timestamp = res[0]["date"].(time.Time)
	} else {
		response.timestamp = time.Now()
	}

	return response, nil

} // fromDatabaseFeed

func FeedHandler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			err     error
			encoded []byte
		)

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(insight)
		if err != nil {
			panic(err)
		}

		response, err := conf.fromDatabaseFeed()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(*response.payload) == 0 {
			if _, err = w.Write([]byte("[]")); err != nil {
				return
			}
			panic(err)
		}

		feedComponents := &feed.Components{
			Endpoint:        r.URL.Path,
			Timestamp:       &response.timestamp,
			Payload:         response.payload,
			Category:        feedCategory,
			ApiEndpoint:     apiEndpoint,
			WebsiteEndpoint: websiteEndpoint,
		}

		if mux.Vars(r)["type"] == "rss" {
			feedData := &rss.Channel{}
			encoded, err = feedData.GenerateFeed(feedComponents)
		} else {
			feedData := &atom.Feed{}
			encoded, err = feedData.GenerateFeed(feedComponents)
		}

		if err != nil {
			panic(err)
		}

		if _, err = w.Write(encoded); err != nil {
			panic(err)
		}

		conf.db.CloseConnection()
	}

} // FeedHandler
