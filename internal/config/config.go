package config

import (
	api "BlogWebApp/internal/http-server"
	db "BlogWebApp/internal/storage/postgres"
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Env                  string `yaml:"env" env-default:"local"`
	db.DBConfig          `yaml:"storage" env-required:"true"`
	api.HTTPServerConfig `yaml:"http_server"`
}

const BasicConfigPath = "./config/local.yaml"

func MustLoad() *Config {
	path := fetchConfigPath()

	if path == "" {
		path = BasicConfigPath
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic("failed to read config" + err.Error())
	}

	return &cfg
}

func fetchConfigPath() string {
	var res string

	// go run main.go --config="./config/local.yaml"
	flag.StringVar(&res, "config", "", "config file path")
	flag.Parse()

	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}

	return res
}
