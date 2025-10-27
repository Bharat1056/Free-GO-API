package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"addr"`
}

// const CONFIG_PATH = "config/local.yaml"

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-default:"production"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server" env-required:"true"`
}

func Mustload() *Config{
	var configPath string

	// here CONFIG_PATH > path the env variable in the system
	// in our case > config/local.yaml
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		// via program argument/flags can we select that
		// flag-name, default-value, description
		flags := flag.String("config", "", "path to the configuration file")
		flag.Parse()

		configPath = *flags // remember here flags gives the pointer

		if configPath == "" {
			log.Fatal("Config path is not set")
		}
	}

	// it checks the configPath is there or not if not is the error comes under os.IsNotExist()
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file doesn't exist: %s", configPath)
	}

	var cfg Config

	// parsing the config file into our struct Config format
	err := cleanenv.ReadConfig(configPath, &cfg)

	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg
}
