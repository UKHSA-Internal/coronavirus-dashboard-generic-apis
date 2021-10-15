package metric_docs

import (
	"fmt"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"generic_apis/apps/utils"
	"generic_apis/db"
	"generic_apis/insight"
	"github.com/gorilla/mux"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type handler struct {
	db          *db.Config
	traceparent string
}

type DocumentationPayload struct {
	LastModified time.Time `json:"last_modified"`
	Body         string    `json:"body"`
}

type Log struct {
	Id      *string `json:"id"`
	Heading *string `json:"heading"`
	Date    *string `json:"date"`
	Expiry  *string `json:"expiry"`
	Type    *string `json:"type"`
}

type Documentation struct {
	Abstract    *DocumentationPayload `json:"abstract,omitempty"`
	Description *DocumentationPayload `json:"description,omitempty"`
	Methodology *DocumentationPayload `json:"methodology,omitempty"`
	Notice      *DocumentationPayload `json:"notice,omitempty"`
	Source      *DocumentationPayload `json:"source,omitempty"`
}

type Payload struct {
	MetricName    string         `json:"metric_name"`
	Metric        string         `json:"metric"`
	Category      string         `json:"category"`
	Documentation *Documentation `json:"documentation"`
	Logs          []*Log         `json:"logs"`
	Tags          []string       `json:"tags"`
}

func stringOrNil(value interface{}) *string {
	if value == nil {
		return nil
	}
	val := value.(string)
	return &val
}

func (conf *handler) fromDatabase(urlParams *map[string]string) (*Payload, error) {

	var (
		params []interface{}
		query  = mainQuery
		pcount = 0
	)

	for key, value := range *urlParams {
		pcount += 1
		query = strings.ReplaceAll(query, fmt.Sprintf(`{%s_token}`, key), strconv.Itoa(pcount))
		params = append(params, value)
	}

	payload := &db.Payload{
		Query:         query,
		Args:          params,
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	res, err := conf.db.FetchAll(payload)
	if err != nil {
		log.Printf(err.Error())
		return nil, fmt.Errorf("failed to retrieve data")
	}

	response := &Payload{}

	if len(res) == 0 {
		return response, nil
	}

	response.MetricName = res[0]["metric_name"].(string)
	response.Metric = res[0]["metric"].(string)
	response.Category = res[0]["category"].(string)
	logs := make([]*Log, len(res[0]["logs"].([]interface{})))
	response.Tags = strings.Split(res[0]["tags"].(string), ",")
	documentations := &Documentation{}

	for index, item := range res[0]["logs"].([]interface{}) {
		logItem := &Log{}
		for key, value := range item.(map[string]interface{}) {
			val := stringOrNil(value)

			switch key {
			case "id":
				logItem.Id = val
				break
			case "heading":
				logItem.Heading = val
				break
			case "date":
				logItem.Date = val
				break
			case "expiry":
				logItem.Expiry = val
				break
			case "type":
				logItem.Type = val
			default:
				continue
			}
		}

		logs[index] = logItem
	}

	// Descending sort
	sort.Slice(logs[:], func(i, j int) bool {
		if strings.Compare(*logs[i].Date, *logs[j].Date) == 1 {
			return true
		} else {
			return false
		}
	})

	response.Logs = logs

	for _, item := range res {
		switch item["asset_type"].(string) {
		case "abstract":
			documentations.Abstract = &DocumentationPayload{
				LastModified: item["last_modified"].(time.Time),
				Body:         item["body"].(string),
			}
			break
		case "description":
			documentations.Description = &DocumentationPayload{
				LastModified: item["last_modified"].(time.Time),
				Body:         item["body"].(string),
			}
			break
		case "methodology":
			documentations.Methodology = &DocumentationPayload{
				LastModified: item["last_modified"].(time.Time),
				Body:         item["body"].(string),
			}
			break
		case "notice":
			documentations.Notice = &DocumentationPayload{
				LastModified: item["last_modified"].(time.Time),
				Body:         item["body"].(string),
			}
			break
		case "source":
			documentations.Source = &DocumentationPayload{
				LastModified: item["last_modified"].(time.Time),
				Body:         item["body"].(string),
			}
			break
		default:
			continue
		}
	}

	response.Documentation = documentations

	return response, nil

} // fromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

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

		pathVars := mux.Vars(r)

		response, err := conf.fromDatabase(&pathVars)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if response.Metric == "" {
			http.NotFound(w, r)
			return
		}

		encoded, err = utils.JSONMarshal(response)
		if err != nil {
			panic(err)
		}

		if _, err = w.Write(encoded); err != nil {
			panic(err)
		}

		conf.db.CloseConnection()
	}

} // Handler
