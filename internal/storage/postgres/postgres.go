package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBConfig struct {
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	Database          string `yaml:"database"`
	User              string `yaml:"user"`
	Password          string `yaml:"password"`
	Migration         bool   `yaml:"migration"`
	Logging           bool   `yaml:"logging"`
	MaxIdleConnection int    `yaml:"max_idle_connection"`
}

var DB *gorm.DB

func (conf *DBConfig) ConnectionString() string { // DSN
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d",
		conf.Host, conf.User, conf.Password, conf.Database, conf.Port,
	)
}

func BuildConnection(c *DBConfig) error {
	gormConfig := gorm.Config{}

	connection, err := gorm.Open(postgres.Open(c.ConnectionString()), &gormConfig)
	if err != nil {
		return fmt.Errorf("failed to load database: %v", err.Error())
	}

	DB = connection

	if c.Migration {
		migrate(DB)
	}

	return nil
}
