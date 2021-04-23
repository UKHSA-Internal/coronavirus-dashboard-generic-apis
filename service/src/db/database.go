package db

import (
	"context"
	"regexp"

	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4"
)

type (
	Config struct {
		DatabaseConnectionString string `env:"POSTGRES_CONNECTION_STRING"`
		database                 *pgx.Conn
	}

	ResultType map[string]interface{}
)

var pattern = regexp.MustCompile(
	`^postgres://(?P<user>[^:]+):?(?P<password>.*)?@(?P<host>.*):?(?P<port>\d*)?/(?P<dbname>.*)?$`)

func (conf *Config) parsedConnStr() *map[string]interface{} {

	if err := env.Parse(conf); err != nil {
		panic(err)
	}

	match := pattern.FindStringSubmatch(conf.DatabaseConnectionString)

	result := make(map[string]interface{})
	for idx, name := range pattern.SubexpNames() {
		if idx > 0 && name != "" {
			result[name] = match[idx]
		}
	}

	if result["port"] == "" {
		result["port"] = uint16(5432)
	}

	return &result

} // parsedConnStr

func (conf *Config) connect() error {

	if err := env.Parse(conf); err != nil {
		panic(err)
	}

	db, err := pgx.Connect(context.Background(), conf.DatabaseConnectionString)
	if err != nil {
		return err
	}

	if err := db.Ping(context.Background()); err != nil {
		return err
	}

	conf.database = db

	return nil

} // connect

func Query(query string, args ...interface{}) ([]ResultType, error) {

	conf := &Config{}
	if err := conf.connect(); err != nil {
		return nil, err
	}
	defer conf.database.Close(context.Background())

	rows, err := conf.database.Query(context.Background(), query, args...)
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

	return results, nil

} // Query
