package configs

import (
	"flag"
	"github.com/pkg/errors"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

const ServiceName = "client"

var options = []option{
	{"config", "string", "", "config file"},

	{"observer.file", "string", "./resources/users.csv", "external users file"},
	{"observer.pause", "int", 5, "observer pause when file doesn't exists (seconds)"},
	{"observer.chunk", "int", 5, "number of parallel requests to service"},
	{"service.target", "string", ":9091", "service target"},
}

type Config struct {
	Observer struct {
		File  string
		Pause int
		Chunk int
	}
	Service struct {
		Target string
	}
}

type option struct {
	name        string
	typing      string
	value       interface{}
	description string
}

// NewConfig returns and prints struct with config parameters
func NewConfig() *Config {
	return &Config{}
}

// read gets parameters from environment variables, flags or file.
func (c *Config) Read() error {
	viper.SetEnvPrefix(ServiceName)
	//viper.SetEnvKeyReplacer(strings.NewReplaces(".", "_"))
	viper.AutomaticEnv()

	for _, o := range options {
		switch o.typing {
		case "string":
			pflag.String(o.name, o.value.(string), o.description)
		case "int":
			pflag.Int(o.name, o.value.(int), o.description)
		case "bool":
			pflag.Bool(o.name, o.value.(bool), o.description)
		case "float64":
			pflag.Float64(o.name, o.value.(float64), o.description)
		default:
			viper.SetDefault(o.name, o.value)
		}
	}

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	viper.BindPFlags(pflag.CommandLine)
	pflag.Parse()

	if fileName := viper.GetString("config"); fileName != "" {
		viper.SetConfigName(fileName)
		viper.SetConfigType("toml")
		viper.AddConfigPath(".")

		if err := viper.ReadInConfig(); err != nil {
			return errors.Wrap(err, "failed to read from file")
		}
	}

	if err := viper.Unmarshal(c); err != nil {
		return errors.Wrap(err, "failed to unmarshal")
	}
	return nil
}
