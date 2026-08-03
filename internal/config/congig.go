package config

import (
	"flag"
	"os"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env string `yaml:"env" env:"Env" env-required:"true"`

	StoragePath string `yaml:"storage_path" env-required:"true"`

	HTTPServer `yaml:"http_server"`
}

func MustLoad() {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flag := flag.String("config", "", "path to the congiguration file")

		flag.Parse()

		configPath = *flag
	}

}
