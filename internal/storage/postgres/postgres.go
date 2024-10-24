package postgres

import (
	"database/sql"
)

type Storage struct {
	DB       *sql.DB
	DBConfig StorageConfig `yaml:"db_config" env-required:"true"`
}

type StorageConfig struct {
	DBHost     string `yaml:"db_host" env-default:"localhost"`
	DBUser     string `yaml:"db_user" env-default:"postgres"`
	DBPassword string `yaml:"db_password" env-required:"true"`
	DBName     string `yaml:"db_name" env-default:"postgres"`
	DBPort     string `yaml:"db_port" env-default:"5432"`
	DBSSLMode  string `yaml:"db_sslmode" env-default:"disabled"`
}

func Connect() {
	// TODO: postgres connection
}
