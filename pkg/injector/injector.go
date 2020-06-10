package injector

import (
	"github.com/CienciaArgentina/go-backend-commons/config"
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
	cfg := config.New(&config.Options{})

	if cfg == nil {
		panic("Config cannot be nil")
	}

	newSQL(cfg)
}

func newSQL(cfg *config.Config) {
	if cfg.Database != nil {
		for _, c := range cfg.Database {
			newDB, name := database.New(&c)
			db[name] = newDB
		}
	}
}
