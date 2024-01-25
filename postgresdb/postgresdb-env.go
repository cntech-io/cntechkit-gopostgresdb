package postgresdb

import (
	"fmt"

	e "github.com/cntech-io/cntechkit-go/v2/env"
	"github.com/joho/godotenv"
)

type PostgresDBEnvName string

const (
	POSTGRESDB_HOST           PostgresDBEnvName = "POSTGRESDB_HOST"
	POSTGRESDB_PORT           PostgresDBEnvName = "POSTGRESDB_PORT"
	POSTGRESDB_DATABASE       PostgresDBEnvName = "POSTGRESDB_DATABASE"
	POSTGRESDB_USERNAME       PostgresDBEnvName = "POSTGRESDB_USERNAME"
	POSTGRESDB_PASSWORD       PostgresDBEnvName = "POSTGRESDB_PASSWORD"
	POSTGRESDB_MIGRATION_FLAG PostgresDBEnvName = "POSTGRESDB_MIGRATION_FLAG"
	POSTGRESDB_SCHEMA         PostgresDBEnvName = "POSTGRESDB_SCHEMA"
)

type postgresDBEnv struct {
	Host          string
	Port          string
	Username      string
	Password      string
	Database      string
	MigrationFlag bool
	Schema        string
}

func NewPostgresDBEnv() *postgresDBEnv {
	if err := godotenv.Load(); err != nil {
		fmt.Println(".env file not found")
	}
	return &postgresDBEnv{
		Host:          e.GetString(string(POSTGRESDB_HOST), false),
		Port:          e.GetString(string(POSTGRESDB_PORT), false),
		Database:      e.GetString(string(POSTGRESDB_DATABASE), false),
		Username:      e.GetString(string(POSTGRESDB_USERNAME), false),
		Password:      e.GetString(string(POSTGRESDB_PASSWORD), false),
		MigrationFlag: e.GetBoolean(string(POSTGRESDB_MIGRATION_FLAG), false),
		Schema:        e.GetString(string(POSTGRESDB_SCHEMA), false),
	}
}
