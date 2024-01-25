package postgresdb

import (
	"fmt"

	e "github.com/cntech-io/cntechkit-go/v2/env"
	"github.com/cntech-io/cntechkit-go/v2/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type postgresDB struct {
	client *gorm.DB
}

var env = NewPostgresDBEnv()

func dsn() string {
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

func NewPostgresDB() *postgresDB {
	return &postgresDB{}
}

func (pdb *postgresDB) Connect() *postgresDB {
	var logMode gormlogger.Interface

	serverEnv := e.NewServerEnv()

	if serverEnv.DebugModeFlag {
		logMode = gormlogger.Default.LogMode(gormlogger.Info)
	} else {
		logMode = gormlogger.Default.LogMode(gormlogger.Silent)
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
	logger.NewLogger(
		&logger.LoggerConfig{AppName: "cntechkit-gopostgresdb"},
	).Info("Connected to PostgreSQL")
	pdb.client = db

	return pdb
}

func (pdb *postgresDB) Migrate(dst ...interface{}) {
	pdb.client.AutoMigrate(dst...)
}

func (pdb *postgresDB) Do() *gorm.DB {
	return pdb.client
}
