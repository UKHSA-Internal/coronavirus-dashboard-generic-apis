package db

import (
	"context"
)

func (conf *Config) fetchAll(query string, args ...interface{}) ([]ResultType, error) {
	rows, err := conf.Database.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	fieldDescriptions := rows.FieldDescriptions()

	var columnNames []string
	for _, col := range fieldDescriptions {
		columnNames = append(columnNames, string(col.Name))
	}
	columnsLen := len(columnNames)

	var results []ResultType

	columns := make([]interface{}, columnsLen)
	columnPointers := make([]interface{}, columnsLen)

	for rows.Next() {
		// Create a slice of interface{}'s to represent each column,
		// and a second slice to contain pointers to each item in the columns slice.
		for index := range columns {
			columnPointers[index] = &columns[index]
		}

		// Scan the result into the column pointers...
		if err := rows.Scan(columnPointers...); err != nil {
			return nil, err
		}

		// Create our map, and retrieve the value for each column from the pointers slice,
		// storing it in the map with the name of the column as the key.
		rowResult := make(ResultType)
		for index, colName := range columnNames {
			val := columnPointers[index].(*interface{})
			rowResult[colName] = *val
		}
		results = append(results, rowResult)
	}

	return results, err
}

func (conf *Config) fetchRow(query string, args ...interface{}) (ResultType, error) {

	rows, err := conf.Database.Query(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	fieldDescriptions := rows.FieldDescriptions()

	var columnNames []string
	for _, col := range fieldDescriptions {
		columnNames = append(columnNames, string(col.Name))
	}
	columnsLen := len(columnNames)

	columns := make([]interface{}, columnsLen)
	columnPointers := make([]interface{}, columnsLen)

	rows.Next()
	// Create a slice of interface{}'s to represent each column,
	// and a second slice to contain pointers to each item in the columns slice.
	for index := range columns {
		columnPointers[index] = &columns[index]
	}

	// Scan the result into the column pointers...
	if err := rows.Scan(columnPointers...); err != nil {
		return nil, err
	}

	// Create our map, and retrieve the value for each column from the pointers slice,
	// storing it in the map with the name of the column as the key.
	rowResult := make(ResultType)
	for index, colName := range columnNames {
		val := columnPointers[index].(*interface{})
		rowResult[colName] = *val
	}

	return rowResult, nil

} // fetchRow
