// Copyright 2020 Luca Stasio. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package core implements core components of gogitsu lib.
//
// controller.go defines the Controller interface and the
// base controller implementation struct from which derive concrete
// controllers into apps.
package core

import (
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

const (

	// MySQLType to create a connection.
	MySQLType = "mysql"
	// PostgresType to create a connection.
	PostgresType = "postgres"
	// Sqlite3Type to create a connection.
	Sqlite3Type = "sqlite3"
	// MSsqlType to create a connection.
	MSsqlType = "mssql"

	// MySQLConnectionStringFormat to create a connection string for MySQL.
	MySQLConnectionStringFormat = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"

	// PostresConnectionStringFormat to create a connection string for Postgres.
	// "host=myhost port=myport user=gorm dbname=gorm password=mypassword"
	PostresConnectionStringFormat = "host=%s port=%d user=%s dbname=%s password=%s"

	// MSsqlConnectionStringFormat to create a connection string for Microsoft SQL Server.
	MSsqlConnectionStringFormat = "sqlserver://%s:%s@%s:%d?database=%s"
)

// DBConfig is the configuration struct to set database info
// for the DBManager initialization.
// Note that the struct fields has gogitsu/conf tags.
type DBConfig struct {
	Type       string `env:"DB_TYPE"`
	Host       string `env:"DB_HOST"`
	Port       int    `env:"DB_PORT"`
	User       string `env:"DB_USER"`
	Password   string `env:"DB_PWD"`
	Database   string `env:"DB_NAME"`
	Log        bool   `env:"DB_LOG_MODE"`
	Migrations struct {
		Enabled            bool `env:"DB_MIG_ENABLED"`
		Drop               bool `env:"DB_MIG_DROP"`
		SingularTableNames bool `env:"DB_MIG_SINGULAR"`
	}
}

// DBManager is the contract intrface for a db manager implementation.
// Each implementation will get a config struct and initialize a db
// connection (pool) accordind to the underline database lib used.
type DBManager interface {
	DB() interface{}
	Configure(config DBConfig)
	Connect() error
}

// DBConfigurator .
type DBConfigurator struct {
	config DBConfig
}

// NewDBConfigurator .
func NewDBConfigurator(config DBConfig) DBConfigurator {
	return DBConfigurator{config: config}
}

// Configure .
func (dbc DBConfigurator) Configure(config DBConfig) {
	dbc.config = config
}

func (dbc DBConfigurator) connectionString() string {
	switch strings.ToLower(dbc.config.Type) {
	case MySQLType:
		return fmt.Sprintf(
			MySQLConnectionStringFormat,
			dbc.config.User,
			dbc.config.Password,
			dbc.config.Host,
			dbc.config.Port,
			dbc.config.Database)
	case PostgresType:
		return fmt.Sprintf(
			PostresConnectionStringFormat,
			dbc.config.Host,
			dbc.config.Port,
			dbc.config.User,
			dbc.config.Database,
			dbc.config.Password)
	case Sqlite3Type:
		return dbc.config.Database
	case MSsqlType:
		return fmt.Sprintf(
			MSsqlConnectionStringFormat,
			dbc.config.User,
			dbc.config.Password,
			dbc.config.Host,
			dbc.config.Port,
			dbc.config.Database)
	default:
		return fmt.Sprintf(
			MySQLConnectionStringFormat,
			dbc.config.User,
			dbc.config.Password,
			dbc.config.Host,
			dbc.config.Port,
			dbc.config.Database)
	}
}

// GormDBManager is the GORM based implementation for DBManager.
type GormDBManager struct {
	DBConfigurator
	db     *gorm.DB
	config DBConfig
}

// NewDefaultGormDBManager return a new empty gorm db manager instance.
func NewDefaultGormDBManager() DBManager {
	return &GormDBManager{DBConfigurator: DBConfigurator{}}
}

// NewGormDBManager return a new gorm db manager instance configured and ready to connect.
func NewGormDBManager(config DBConfig) DBManager {
	return &GormDBManager{DBConfigurator: NewDBConfigurator(config)}
}

// DB return the internal gorm DB.
func (gbm *GormDBManager) DB() interface{} {
	return gbm.db
}

// Connect establish a connection (pool) to the db.
func (gbm *GormDBManager) Connect() error {
	connectionString := gbm.connectionString()
	//dbc.logger.Debugf("Connecting to DB at %s", connectionString)

	var err error
	gbm.db, err = gorm.Open(gbm.config.Type, connectionString)
	if err != nil {
		return err
	}

	gbm.db.LogMode(gbm.config.Log)

	return nil
}
