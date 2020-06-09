package config

import (
	"fmt"
	"github.com/CienciaArgentina/go-backend-commons/pkg/scope"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	ApiKey   string     `yaml:"apikey"`
	Database []Database `yaml:"database"`
	Server   Server     `yaml:"server"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"` // This is the name of the environment variable, not the actual value
	Hostname string `yaml:"hostname"` // This is the name of the environment variable, not the actual value
	Database string `yaml:"database"`
}

type Server struct {
	Port string `yaml:"server_port"`
}

type Options struct {
	FilePath string
	Scope    string
	IsCloud  bool
}

func New(o ...*Options) *Config {
	c := &Config{}

	if o == nil {
		o = []*Options{}
	}

	if o[0] == nil {
		o[0] = SetDefaultOptions()
	}

	if o[0].Scope == "" {
		if scope.IsCloud() {
			if scope.IsProductiveScope() {
				o[0].Scope = scope.GetScope()
			} else {
				o[0].Scope = scope.Development
			}
		} else {
			o[0].Scope = scope.Local
		}
	}

	var data []byte
	var err error

	if o[0].FilePath == "" {
		data, err = ioutil.ReadFile(fmt.Sprintf("./config/config.%s.yml", o[0].Scope))
	} else {
		data, err = ioutil.ReadFile(o[0].FilePath)
	}

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(data, &c)
	if err != nil {
		panic(err)
	}

	return c
}

func SetDefaultOptions() *Options {
	return &Options{
		FilePath: "",
		Scope:    scope.Local,
		IsCloud:  false,
	}
}
