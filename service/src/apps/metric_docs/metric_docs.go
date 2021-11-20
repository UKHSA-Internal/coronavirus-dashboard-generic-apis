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
	Id           *string   `json:"id"`
	Heading      *string   `json:"heading"`
	Date         *string   `json:"date"`
	Expiry       *string   `json:"expiry"`
	Type         *string   `json:"type"`
	ApplicableTo *[]string `json:"applicable_to"`
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
	Deprecated    *string        `json:"deprecated"`
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
	response.Deprecated = stringOrNil(res[0]["deprecated"])
	logs := make([]*Log, len(res[0]["logs"].([]interface{})))
	response.Tags = strings.Split(res[0]["tags"].(string), ",")
	documentations := &Documentation{}

	// Same metric, same logs - first item will suffice.
	for index, item := range res[0]["logs"].([]interface{}) {
		logItem := &Log{}
		for key, value := range item.(map[string]interface{}) {
			switch key {
			case "id":
				logItem.Id = stringOrNil(value)
				break
			case "heading":
				logItem.Heading = stringOrNil(value)
				break
			case "date":
				logItem.Date = stringOrNil(value)
				break
			case "expiry":
				logItem.Expiry = stringOrNil(value)
				break
			case "type":
				logItem.Type = stringOrNil(value)
				break
			case "applicable_to":

				// Using maps to ensure uniqueness.
				areas := make(map[string]bool, len(value.([]interface{})))

				for _, area := range value.([]interface{}) {
					switch area {
					case "overview::^K.*$":
						areas["UK"] = true
						continue
					case "nation::^E92000001$":
						areas["England"] = true
						continue
					case "region::^E.*$":
						areas["England regions"] = true
						continue
					case "utla::^E.*$":
						areas["England UTLAs"] = true
						continue
					case "ltla::^E.*$":
						areas["England LTLAs"] = true
						continue
					case "msoa::^E.*$":
						areas["England MSOAs"] = true
						continue
					case "nation::^S92000003$":
						areas["Scotland"] = true
						continue
					case "utla::^S.*$":
					case "ltla::^S.*$":
						areas["Scotland local authorities"] = true
						continue
					case "nation::^N92000002$":
						areas["Northern Ireland"] = true
						continue
					case "utla::^N.*$":
					case "ltla::^N.*$":
						areas["Northern Ireland local authorities"] = true
						continue
					case "nation::^W.*$":
						areas["Wales"] = true
						continue
					case "utla::^W.*$":
					case "ltla::^W.*$":
						areas["Wales local authorities"] = true
						continue
					case "nhsRegion::^.*$":
						areas["All NHS regions"] = true
						continue
					case "nhsTrust::^.*$":
						areas["All NHS trusts"] = true
						continue
					default:
						areas["Unspecified"] = true
						continue
					}
				}

				uniqueAreas := make([]string, len(areas))
				areaInd := 0
				for areaItem := range areas {
					// Deduplicated items leave empty spaces
					// in the map.
					if areaItem != "" {
						uniqueAreas[areaInd] = areaItem
						areaInd++
					}
				}

				logItem.ApplicableTo = &uniqueAreas

			default:
				continue
			}
		}

		logs[index] = logItem
	}

	// Descending sort
	sort.Slice(logs[:], func(i, j int) bool {
		return strings.Compare(*logs[i].Date, *logs[j].Date) == 1
	})

	response.Logs = logs

	for _, item := range res {
		if item["asset_type"] == nil {
			continue
		}

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
