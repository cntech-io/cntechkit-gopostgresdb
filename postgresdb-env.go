package cntechkitgopostgresdb

import (
	"fmt"

	"github.com/cntech-io/cntechkit-go/utils"
	"github.com/joho/godotenv"
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
		Host:          utils.GetStringEnv(string(POSTGRESDB_HOST), false),
		Port:          utils.GetStringEnv(string(POSTGRESDB_PORT), false),
		Database:      utils.GetStringEnv(string(POSTGRESDB_DATABASE), false),
		Username:      utils.GetStringEnv(string(POSTGRESDB_USERNAME), false),
		Password:      utils.GetStringEnv(string(POSTGRESDB_PASSWORD), false),
		MigrationFlag: utils.GetBooleanEnv(string(POSTGRESDB_MIGRATION_FLAG), false),
		Schema:        utils.GetStringEnv(string(POSTGRESDB_SCHEMA), false),
	}
}
