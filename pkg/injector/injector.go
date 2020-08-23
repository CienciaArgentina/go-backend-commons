package injector

import (
	"errors"
	"github.com/CienciaArgentina/go-backend-commons/config"
	"github.com/CienciaArgentina/go-backend-commons/pkg/clog"
	"github.com/CienciaArgentina/go-backend-commons/pkg/infrastructure/database"
)

var (
	db = make(map[string]*database.Database)
)

func GetDB(name string) *database.Database {
	name = database.FormatDbName(name)
	return db[name]
}

func Initilize() {
	cfg, err := config.New()

	if err != nil {
		return
	}

	if cfg == nil {
		msg := "Config cannot be nil"
		clog.Panic(msg, "initialize-injector", errors.New(msg), nil)
		return
	}

	newSQL(cfg)
}

func newSQL(cfg *config.Config) {
	if cfg.Database != nil {
			newDB, name := database.New(cfg.Database)
			db[name] = newDB
		}
}
