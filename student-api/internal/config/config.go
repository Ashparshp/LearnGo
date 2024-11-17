package config

import (
	"flag"
	"os"
	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env string `yaml:"env" env:"ENV" emv-required:"true"`
	StoragePath  string `yaml:"storage_path" env-required:"true"`
	HTTPServer HTTPServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")
	if configPath == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		configPath = *flags


		if configPath == "" {
			panic("config path is required")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist")
	}

	var cfg Config
	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}