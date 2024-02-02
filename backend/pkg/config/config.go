package config

import (
	"github.com/spf13/viper"
	"log"
)

// Config to map the config vars.
type Config struct {
	DBUser           string `mapstructure:"DBUSER"`
	DBPass           string `mapstructure:"DBPASS"`
	DBIp             string `mapstructure:"DBIP"`
	DBName           string `mapstructure:"DBNAME"`
	Port             string `mapstructure:"PORT"`
	JwtSecret        string `mapstructure:"JWT_SECRET"`
	JwtExpireMinutes int    `mapstructure:"JWT_EXPIRE_MINUTES"`
}

// InitConfig reads the app.env file and maps the config vars.
func InitConfig() *Config {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	//automatically reads the config vars
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}
	var config *Config

	//converts the read config vars into mapped struct type
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal("Error reading env file", err)
	}
	return config
}

// LocalConfig holds the config vars.
var LocalConfig *Config

// SetConfig sets the config vars.
func SetConfig() {
	LocalConfig = InitConfig()
}
