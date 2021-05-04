package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

// dbResponse holds response and metadata for a database
// query.
type dbResponse struct {
	// columnNames holds name of columns included in a
	// database response.
	columnNames []string

	// rows holds the rows returned in response to a
	// query from the database.
	rows pgx.Rows

	// columnsLen holds the number of columns in a
	// database response.
	columnsLen int
}

// fetch makes a call to the the database and prepares
// the response for additional processing.
func (conf *Config) fetch(query string, args ...interface{}) (*dbResponse, error) {
	var (
		response = &dbResponse{}
		err      error
	)

	response.rows, err = conf.Database.Query(context.Background(), query, args...)
	if err != nil {
		return response, err
	}

	fieldDescriptions := response.rows.FieldDescriptions()

	for _, col := range fieldDescriptions {
		response.columnNames = append(response.columnNames, string(col.Name))
	}

	response.columnsLen = len(response.columnNames)

	return response, nil

} // fetch

// fetchAll retrieves and prepares all data returned in
// response to a query.
func (conf *Config) fetchAll(query string, args ...interface{}) ([]ResultType, error) {

	var (
		results  []ResultType
		response *dbResponse
		err      error
	)

	response, err = conf.fetch(query, args...)
	if err != nil {
		return nil, err
	}

	columns := make([]interface{}, response.columnsLen)
	columnPointers := make([]interface{}, response.columnsLen)

	for response.rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		for index := range columns {
			columnPointers[index] = &columns[index]
		}

		// Scan the result into the column pointers...
		if err = response.rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		rowResult := make(ResultType)
		for index, colName := range response.columnNames {
			val := columnPointers[index].(*interface{})
			rowResult[colName] = *val
		}
		results = append(results, rowResult)
	}

	return results, err

} // fetchAll

// fetchRow retrieves and prepares the first row of data
// returned in response to a query.
func (conf *Config) fetchRow(query string, args ...interface{}) (ResultType, error) {

	var (
		response *dbResponse
		err      error
	)

	response, err = conf.fetch(query, args...)
	if err != nil {
		return nil, err
	}

	columns := make([]interface{}, response.columnsLen)
	columnPointers := make([]interface{}, response.columnsLen)

	response.rows.Next()
	// Create a slice of interface{}'s to represent each column,
	// and a second slice to contain pointers to each item in the columns slice.
	for index := range columns {
		columnPointers[index] = &columns[index]
	}

	// Scan the result into the column pointers...
	if err = response.rows.Scan(columnPointers...); err != nil {
		return nil, err
	}

	// Create our map, and retrieve the value for each column from the pointers slice,
	// storing it in the map with the name of the column as the key.
	rowResult := make(ResultType)
	for index, colName := range response.columnNames {
		val := columnPointers[index].(*interface{})
		rowResult[colName] = *val
	}

	return rowResult, nil

} // fetchRow
