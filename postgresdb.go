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
)

type postgresDBKit struct {
	client *gorm.DB
}

func dsn() string {
	env := NewPostgresDBEnv()
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
	var db *gorm.DB
	var err error
	var logMode logger.Interface

	serverEnv := gokit.NewServerEnv()

	if serverEnv.DebugModeFlag {
		logMode = logger.Default.LogMode(logger.Info)
	} else {
		logMode = logger.Default.LogMode(logger.Silent)
	}

	db, err = gorm.Open(postgres.Open(dsn()), &gorm.Config{
		QueryFields: true,
	}, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "cntech.", // schema name
			SingularTable: false,
		},
		Logger: logMode,
	})
	if err != nil {
		panic("Failed to connect to PostgreSQL: " + err.Error())
	}
	pdb.client = db

	return pdb
}

func (pdb *postgresDBKit) Migrate(dst ...interface{}) {
	pdb.client.AutoMigrate(dst...)
}

func (pdb *postgresDBKit) Do() *gorm.DB {
	return pdb.client
}
