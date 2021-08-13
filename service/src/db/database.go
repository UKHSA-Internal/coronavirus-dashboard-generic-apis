package db

import (
	"context"
	"encoding/json"
	"time"

	"generic_apis/insight"
	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4"
	"github.com/microsoft/ApplicationInsights-Go/appinsights"
)

type (
	Config struct {
		DatabaseConnectionString string `env:"POSTGRES_CONNECTION_STRING"`
		Database                 *pgx.Conn
		dbConfig                 *pgx.ConnConfig
		Insight                  appinsights.TelemetryClient
	}

	Payload struct {
		Query         string
		Args          []interface{}
		OperationData *insight.OperationData
	}

	ResultType map[string]interface{}
)

func (conf *Config) CloseConnection() {

	err := conf.Database.Close(context.Background())

	if err != nil {
		panic(err)
	}

} // closeConnection

// Connect establishes a connection to the DB.
func Connect(insight appinsights.TelemetryClient) (*Config, error) {

	var err error
	conf := &Config{Insight: insight}
	if err = env.Parse(conf); err != nil {
		panic(err)
	}

	conf.dbConfig, err = pgx.ParseConfig(conf.DatabaseConnectionString + "?client_encoding=utf8")
	if err != nil {
		return nil, err
	}

	// Support for PGBouncer
	conf.dbConfig.BuildStatementCache = nil
	conf.dbConfig.PreferSimpleProtocol = true
	conf.dbConfig.ConnectTimeout = 10 * time.Second

	conf.Database, err = pgx.ConnectConfig(context.Background(), conf.dbConfig)
	if err != nil {
		return nil, err
	}

	if err = conf.Database.Ping(context.Background()); err != nil {
		return nil, err
	}

	return conf, nil

} // Connect

func (conf *Config) trackRequest(payload *Payload, action string) func(error) {

	startTime := time.Now()

	return func(err error) {
		endTime := time.Now()

		dependency := appinsights.NewRemoteDependencyTelemetry(
			"database",
			"postgresql",
			"database",
			err == nil)

		dependency.Data = payload.Query
		dependency.MarkTime(startTime, endTime)

		argsJson, jsonErr := json.Marshal(payload.Args)
		if jsonErr == nil {
			dependency.Properties["args"] = string(argsJson)
		}

		dependency.Properties["action"] = action

		dependency.Id = insight.GenerateOperationId()
		dependency.Tags.Operation().SetParentId(payload.OperationData.ParentId)
		dependency.Tags.Operation().SetId(payload.OperationData.OperationId)
		dependency.Tags.Cloud().SetRole(payload.OperationData.CloudRoleName)
		dependency.Tags.Cloud().SetRoleInstance(payload.OperationData.CloudRoleInstance)

		conf.Insight.Track(dependency)

	}

} // trackRequest

// FetchAll establishes a connection to the Database, executes `query`
// with `args` and returns all returning rows.
func (conf *Config) FetchAll(payload *Payload) ([]ResultType, error) {

	trackerDone := conf.trackRequest(payload, "FetchAll")
	response, responseErr := conf.fetchAll(payload.Query, payload.Args...)
	trackerDone(responseErr)

	return response, responseErr

} // FetchAll

// FetchRow establishes a connection to the Database, executes `query`
// with `args` and returns the first rows.
func (conf *Config) FetchRow(payload *Payload) (ResultType, error) {

	trackerDone := conf.trackRequest(payload, "FetchRow")
	response, responseErr := conf.fetchRow(payload.Query, payload.Args...)
	trackerDone(responseErr)

	return response, responseErr

} // FetchRow
