package config

import (
	"errors"
	"github.com/CienciaArgentina/go-backend-commons/pkg/clog"
	"os"
)

const (
	envDBUsername = "DB_USERNAME"
	envDBPassword = "DB_PASSWORD"
	envDBHostname = "DB_HOSTNAME"
	EnvDBName = "DB_NAME"
	envDBPort = "DB_PORT"
)

type Config struct {
	Database *Database
}

type Database struct {
	Username string
	Password string
	Hostname string
	Database string
	Port string
}

type Options struct {
	FilePath string
	Scope    string
	IsCloud  bool
}

func New() (*Config, error) {
	db, err := getDBEnv()
	if err != nil {
		clog.Panic("Error with database config", "new-config", err, nil)
		return nil, err
	}

	cfg := &Config{Database: db}

	return cfg, nil
}

func getDBEnv() (*Database, error) {
	username := os.Getenv(envDBUsername)
	password := os.Getenv(envDBPassword)
	hostname := os.Getenv(envDBHostname)
	name := os.Getenv(EnvDBName)
	port := os.Getenv(envDBPort)
	db := &Database{
		Username: username,
		Password: password,
		Hostname: hostname,
		Database: name,
		Port:     port,
	}

	return db, db.IsValid()
}

func (d *Database) IsValid() error {
	if d.Username == "" {
		return errors.New("missing database username")
	}

	if d.Password == "" {
		return errors.New("missing database password")
	}

	if d.Hostname == "" {
		return errors.New("missing database hostname")
	}

	if d.Database == "" {
		return errors.New("missing database name")
	}

	if d.Port == "" {
		return errors.New("missing database port")
	}

	return nil
}

