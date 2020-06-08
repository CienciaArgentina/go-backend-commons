package config

type Config struct {
	Database []Database `yaml:"database"`
	Server Server `yaml:"server"`
}

type Database struct {
	Username string `yaml:"username"`
	Password string `envconfig:"ENV_DB_PASSWORD"`
	Hostname string `envconfig:"ENV_DB_HOSTNAME"`
	Port     string `yaml:"port"`
	Database string `envconfig:"ENV_DB_NAME"`
}

type Server struct {
	Port string `yaml:"server_port"`
}

type Options struct {
	FilePath string
	Scope string
	IsCloud bool
}

func New(o ...*Options) *Config {
	c := &Config{}

	//if o == nil || o[0] == nil {
	//	o := SetDefaultOptions()
	//}

	return c
}

func SetDefaultOptions() *Options {
	return &Options{
		FilePath: "",
		Scope:    "",
		IsCloud:  false,
	}
}