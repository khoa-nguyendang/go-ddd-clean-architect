package config

import (
	"errors"
	"log"

	"github.com/spf13/viper"
)

// App config struct
type Config struct {
	Server     ServerConfig
	Mysql      MysqlConfig
	Logger     Logger
	Kafka      Kafka
	OpenSearch OpenSearch
}

// Server config struct
type ServerConfig struct {
	AppVersion string
	Port       string
	Debug      bool
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// MySQL config
type MysqlConfig struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUser     string
	MysqlPassword string
	MysqlDbname   string
	MysqlDriver   string
}

// Kafka config
type Kafka struct {
	Server         string
	AuditLogServer string
	Username       string
	Password       string
}

// OpenSearch config
type OpenSearch struct {
	Server   string
	Username string
	Password string
}

// Load config file from given path
func LoadConfig(filename string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(filename)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}

// Get config
func GetConfig(configPath string) (*Config, error) {
	cfgFile, err := LoadConfig(configPath)
	if err != nil {
		return nil, err
	}

	cfg, err := ParseConfig(cfgFile)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config.dev"
}
