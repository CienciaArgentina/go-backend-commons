package database

import (
	"fmt"
	"github.com/CienciaArgentina/go-backend-commons/config"
	"github.com/CienciaArgentina/go-backend-commons/pkg/scope"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"os"
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

	if cfg.Username == "" || cfg.Database == "" || cfg.Hostname == "" || cfg.Password == "" {
		panic("Invalid DB config")
	}

	dbName := FormatDbName(cfg.Database)

	password := os.Getenv(cfg.Password)
	hostname := os.Getenv(cfg.Hostname)

	if password == "" || hostname == "" {
		panic("There's no environment variable with a valid password or hostname")
	}

	db, _ := sqlx.Open(MySQL, fmt.Sprintf("%s:%s@(%s)/%s", cfg.Username, cfg.Password, cfg.Hostname, cfg.Database))

	if err := db.Ping(); err != nil {
		panic(err)
	}

	return &Database{Database: db}, dbName
}

func FormatDbName(dbname string) string {
	if suffix := string(dbname[len(dbname)-2]); suffix == "dev" && !scope.IsProductiveScope() {
		return strings.TrimSuffix(dbname, suffix)
	}
	return dbname
}
