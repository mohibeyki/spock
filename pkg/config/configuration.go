package config

import (
	"log"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/spf13/viper"
)

// Config is the global configuration data
var Config *Configuration

// Configuration struct holds all configuration data types
type Configuration struct {
	Auth     AuthConfiguration
	Server   ServerConfiguration
	Database DatabaseConfiguration
}

// Init initializes configuration
func Init() {
	var configuration *Configuration

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	configuration.Auth.Algorithm = jwt.NewHS512([]byte(configuration.Auth.PrivateKey))
	Config = configuration
}

// GetConfig returns configuration data
func GetConfig() *Configuration {
	return Config
}
