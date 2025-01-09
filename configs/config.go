package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var cfg *config

type config struct {
	API   APIConfig
	DB    DBConfig
	JWT   JWTConfig
	ADMIN AdminConfig
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

type JWTConfig struct {
	SecretKey string
}

type AdminConfig struct {
	UToken string
}

func GetConfig() *config {
	return cfg
}

func init() {
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
		JWT: JWTConfig{
			SecretKey: viper.GetString("jwt.secret_key"),
		},
		ADMIN: AdminConfig{
			UToken: viper.GetString("admin.id"),
		},
	}
	if os.Getenv("GINMODE") == "local" {
		cfg = &config{
			API: APIConfig{
				Port: viper.GetString("local_api.port"),
			},
			DB: DBConfig{
				Host:     viper.GetString("local_db.host"),
				Port:     viper.GetInt("local_db.port"),
				User:     viper.GetString("local_db.user"),
				Password: viper.GetString("local_db.password"),
				DBName:   viper.GetString("local_db.dbname"),
				SSLmode:  viper.GetString("local_db.sslmode"),
			},
			JWT: JWTConfig{
				SecretKey: viper.GetString("local_jwt.secret_key"),
			},
			ADMIN: AdminConfig{
				UToken: viper.GetString("local_admin.id"),
			},
		}
	}
}
