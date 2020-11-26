package config

import (
	"time"

	"github.com/spf13/viper"
)

// AppConfig ...
type AppConfig struct {
	// Common
	AppName  string
	HttpPort int
	GrpcPort int

	// Database
	Database Database

	// Api clients
	ApiClients map[string]ApiClient
}

// Database ...
type Database struct {
	Addr     string
	Database string
	User     string
	Password string
	Options  DatabaseOptions
}

// DatabaseOptions ...
type DatabaseOptions struct {
	PoolSize           int
	MinIdleConns       int
	MaxConnAge         time.Duration
	PoolTimeout        time.Duration
	IdleTimeout        time.Duration
	IdleCheckFrequency time.Duration
	LogQuery           bool
}

// ApiClient ...
type ApiClient struct {
	Target  string
	Options ApiClientOptions
}

// ApiClientOptions ...
type ApiClientOptions struct {
	PoolSize   int
	RetryCount int
	Timeout    time.Duration
}

var appConfig *AppConfig

// GetCfg ...
func GetCfg() *AppConfig {
	return appConfig
}

// InitConfigFromFile ...
func InitConfigFromFile(configFile string) (*AppConfig, error) {
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	appCfg := &AppConfig{}
	err = viper.Unmarshal(appCfg)
	appConfig = appCfg
	return appConfig, err
}
