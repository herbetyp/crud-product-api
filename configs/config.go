package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
	SSLmode  string
}

func GetConfig() *config {
	return cfg
}

func init() {
	viper.SetDefault("api.port", "3000")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")
	viper.SetDefault("db.sslmode", "disable")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("toml")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			fmt.Printf("Error reading config file: %v\n", err)
			return
		}
	}

	cfg = &config{
		API: APIConfig{
			Port: viper.GetString("api.port"),
		},
		DB: DBConfig{
			Host:     viper.GetString("db.host"),
			Port:     viper.GetInt("db.port"),
			User:     viper.GetString("db.user"),
			Password: viper.GetString("db.password"),
			DBName:   viper.GetString("db.dbname"),
			SSLmode:  viper.GetString("db.sslmode"),
		},
	}
}
