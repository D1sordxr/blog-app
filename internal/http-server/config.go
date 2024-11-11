package http_server

import "time"

type HTTPServerConfig struct {
	Address     string        `yaml:"address" env-required:"true"`
	Port        string        `yaml:"port" env-required:"true"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}
