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
	"search": `[a-zA-Z0-9,-]`,
	"page":   `\d{1,3}`,
	"title":  `[a-z]{4,120}`,
	"type":   `[a-z]{5,40}`,
}

// TODO:
// 	Add month query

func (conf *handler) fromDatabase(date string, queryParams url.Values) ([]db.ResultType, error) {

	var (
		params  []interface{}
		filters []string
		query   = simpleQuery
		pcount  = 0
	)

	if date != "" {
		pcount += 1
		filters = append(filters, strings.ReplaceAll(queryParamFilters["date"], "{token_id}", strconv.Itoa(pcount)))
		date += "-01"
		params = append(params, date)
	}

	if _, ok := queryParams["page"]; !ok {
		// Set default page
		queryParams["page"] = []string{"1"}
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
			pageValue, err := strconv.Atoi(value)
			if err != nil {
				return nil, err
			}
			pageValue -= 1
			// Page limit is 20 (defined in pagination query).
			query = strings.Replace(query, PaginationToken, fmt.Sprintf(paginationQuery, pageValue*20), 1)
		} else {
			if _, ok := queryParams["search"]; ok {
				// Search queries use a different base query.
				pcount += 1
				query = strings.ReplaceAll(searchQuery, QueryToken, strconv.Itoa(pcount))
			}
			pcount += 1
			params = append(params, value)
			filters = append(filters, strings.ReplaceAll(queryParamFilters[key], QueryToken, strconv.Itoa(pcount)))
		}
	}

	joinedFilters := strings.Join(filters, " AND ")
	if joinedFilters != "" {
		joinedFilters = fmt.Sprintf(filtersQuery, joinedFilters)
	}
	query = strings.Replace(query, FiltersToken, joinedFilters, 1)

	payload := &db.Payload{
		Query:         query,
		Args:          params,
		OperationData: insight.GetOperationData(conf.traceparent),
	}

	res, err := conf.db.FetchAll(payload)

	return res, err

} // FromDatabase

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

		response, err := conf.fromDatabase(pathVars["date"], r.URL.Query())
		if err != nil {
			http.Error(w, "failed to retrieve data", http.StatusBadRequest)
		}

		if len(response) == 0 {
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

} // queryByCode
