package cntechkitgopostgresdb

import (
	"fmt"

	gokit "github.com/cntech-io/cntechkit-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
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

type postgresDBKit struct {
	client *gorm.DB
}

func dsn() string {
	env := NewPostgresDBEnv()
	if env.Host == "" {
		panic("POSTGRESDB_HOST is not set")
	}
	if env.Port == "" {
		panic("POSTGRESDB_PORT is not set")
	}
	if env.Database == "" {
		panic("POSTGRESDB_DATABASE is not set")
	}
	if env.Username == "" {
		panic("POSTGRESDB_USERNAME is not set")
	}
	if env.Password == "" {
		panic("POSTGRESDB_PASSWORD is not set")
	}
	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable",
		env.Host,
		env.Username,
		env.Password,
		env.Database,
		env.Port,
	)
}

func NewPostgresDB() *postgresDBKit {
	return &postgresDBKit{}
}

func (pdb *postgresDBKit) Connect() *postgresDBKit {
	var logMode logger.Interface

	serverEnv := gokit.NewServerEnv()
	env := NewPostgresDBEnv()

	if serverEnv.DebugModeFlag {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	db, err := gorm.Open(postgres.Open(dsn()), &gorm.Config{
		QueryFields: true,
	}, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%v.", env.Schema),
			SingularTable: false,
		},
		Logger: logMode,
	})
	if err != nil {
		panic("Failed to connect to PostgreSQL: " + err.Error())
	}
	gokit.NewLogger(
		&gokit.LoggerConfig{AppName: "cntechkit-gopostgresdb"},
	).Info("Connected to PostgreSQL")
	pdb.client = db

	return pdb
}

func (pdb *postgresDBKit) Migrate(dst ...interface{}) {
	pdb.client.AutoMigrate(dst...)
}

func (pdb *postgresDBKit) Do() *gorm.DB {
	return pdb.client
}
