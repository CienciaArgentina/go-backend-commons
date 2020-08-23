package database

import (
	"errors"
	"fmt"
	"github.com/CienciaArgentina/go-backend-commons/config"
	"github.com/CienciaArgentina/go-backend-commons/pkg/clog"
	"github.com/CienciaArgentina/go-backend-commons/pkg/scope"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"strings"
)

const (
	MySQL = "mysql"
)

type Database struct {
	Database *sqlx.DB
	Mock     sqlmock.Sqlmock
}

func New(cfg *config.Database) (*Database, string) {
	if scope.IsTesting() {
		db, mock, _ := sqlmock.New()
		defer db.Close()
		sqlxDb := sqlx.NewDb(db, "sqlmock")
		return &Database{Database: sqlxDb, Mock: mock}, cfg.Database
	}

	dbName := FormatDbName(cfg.Database)

	db, _ := sqlx.Open(MySQL, fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.Username, cfg.Password, cfg.Hostname, cfg.Port, cfg.Database))

	if err := db.Ping(); err != nil {
		msg := "error pinging db"
		clog.Panic(msg, "new-mysql", errors.New(msg), nil)
		return nil, ""
	}

	return &Database{Database: db}, dbName
}

func FormatDbName(dbname string) string {
	if suffix := string(dbname[len(dbname)-2]); suffix == "dev" && !scope.IsProductiveScope() {
		return strings.TrimSuffix(dbname, suffix)
	}
	return dbname
}
