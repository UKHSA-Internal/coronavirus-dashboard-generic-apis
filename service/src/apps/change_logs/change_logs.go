package change_logs

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

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

var paramPatterns = map[string]string{
	"search": `[a-zA-Z0-9,'\s-]{4,40}`,
	"page":   `\d{1,3}`,
	"title":  `[a-zA-Z:\s-]{4,120}`,
	"type":   `[a-zA-Z\s]{5,40}`,
}

const pageLimit = 20

func (conf *handler) fromDatabase(date, id string, queryParams url.Values) (db.ResultType, error) {

	var (
		err     error
		params  []interface{}
		filters = []string{releaseFilter, metricFilter}
		query   = simpleQuery
		pcount  = 0
		page    = 1
	)

	if date != "" {
		pcount += 1
		filters = append(filters, strings.ReplaceAll(queryParamFilters["date"], queryToken, strconv.Itoa(pcount)))
		date += "-01"
		params = append(params, date)
	}

	if _, ok := queryParams["page"]; !ok {
		// Set default page
		queryParams["page"] = []string{"1"}
	}

	if id != "" {
		pcount += 1
		filters = []string{
			strings.ReplaceAll(queryParamFilters["record"], queryToken, strconv.Itoa(pcount)),
		}
		query = recordQuery
		params = append(params, id)
	}

	for key, pattern := range paramPatterns {
		value := queryParams.Get(key)
		if value == "" {
			continue
		}

		if !utils.ValidateParam(pattern, value) {
			return nil, fmt.Errorf("invalid query param")
		}

		if key == "page" {
			if value == "" {
				value = "1"
			}
			page, err = strconv.Atoi(value)
			if err != nil {
				panic(err)
			}

			offset := (page - 1) * pageLimit

			// Page limit is 20 (defined in pagination query).
			query = strings.Replace(
				query,
				paginationToken,
				fmt.Sprintf(paginationQuery, offset),
				1,
			)
		} else {
			if _, ok := queryParams["search"]; ok {
				// Search queries use a different base query.
				pcount += 1
				query = strings.ReplaceAll(searchQuery, queryToken, strconv.Itoa(pcount))
			}
			pcount += 1
			params = append(params, value)
			filters = append(filters, strings.ReplaceAll(queryParamFilters[key], queryToken, strconv.Itoa(pcount)))
		}
	}

	joinedFilters := strings.Join(filters, " AND ")
	if joinedFilters != "" {
		joinedFilters = fmt.Sprintf(filtersQuery, joinedFilters)
	}
	query = strings.Replace(query, filtersToken, joinedFilters, 1)

	payload := &db.Payload{
		Query:         query,
		Args:          params,
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	res, err := conf.db.FetchRow(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve data")
	}

	res["page"] = page

	if res["data"] != nil {
		res["length"] = len(res["data"].([]interface{}))
	} else {
		res["data"] = []interface{}{}
	}

	return res, nil

} // fromDatabase

func Handler(insight appinsights.TelemetryClient) func(w http.ResponseWriter, r *http.Request) {

	conf := &handler{}

	return func(w http.ResponseWriter, r *http.Request) {

		var (
			isComponentQuery bool
			err              error
			response         interface{}
			encoded          []byte
			componentName    string
		)

		conf.traceparent = r.Header.Get("traceparent")

		conf.db, err = db.Connect(insight)
		if err != nil {
			panic(err)
		}

		pathVars := mux.Vars(r)
		componentName, isComponentQuery = pathVars["component"]

		if isComponentQuery {
			response, err = conf.getComponentsFromDatabase(componentName)
		} else {
			response, err = conf.fromDatabase(pathVars["date"], pathVars["id"], r.URL.Query())
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Return 204 if payload is empty.
		if !isComponentQuery {
			lenKeys := len(response.(db.ResultType))
			lenData := len(response.(db.ResultType)["data"].([]interface{}))

			if lenKeys == 0 || lenData == 0 {
				w.WriteHeader(http.StatusNoContent)
				return
			} else if _, isIdQuery := pathVars["id"]; isIdQuery && lenData == 1 {
				response = response.(db.ResultType)["data"].([]interface{})[0]
			}
		} else if len(response.([]db.ResultType)) == 0 {
			w.WriteHeader(http.StatusNoContent)
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
