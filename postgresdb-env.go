package cntechkitgopostgresdb

import (
	"github.com/cntech-io/cntechkit-go/utils"
)

type postgresDBEnv struct {
	Host          string
	Port          string
	Username      string
	Password      string
	Database      string
	MigrationFlag bool
}

func NewPostgresDBEnv() *postgresDBEnv {
	return &postgresDBEnv{
		Host:          utils.GetStringEnv(string(POSTGRESDB_HOST), false),
		Port:          utils.GetStringEnv(string(POSTGRESDB_PORT), false),
		Database:      utils.GetStringEnv(string(POSTGRESDB_DATABASE), false),
		Username:      utils.GetStringEnv(string(POSTGRESDB_USERNAME), false),
		Password:      utils.GetStringEnv(string(POSTGRESDB_PASSWORD), false),
		MigrationFlag: utils.GetBooleanEnv(string(POSTGRESDB_MIGRATION_FLAG), false),
	}
}
