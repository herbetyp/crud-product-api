package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

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
	Host               string
	Port               int
	User               string
	Password           string
	DBName             string
	SSLmode            string
	SetMaxIdleConns    int
	SetMaxOpenConns    int
	SetConnMaxLifetime time.Duration
}

type JWTConfig struct {
	SecretKey string
	ExpiresIn time.Duration
}

type AdminConfig struct {
	UId string
}

func InitConfig() {
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

	switch os.Getenv("GINMODE") {
	case "release":
		cfg = &config{
			API: APIConfig{
				Port: viper.GetString("api.port"),
			},
			DB: DBConfig{
				Host:               viper.GetString("db.host"),
				Port:               viper.GetInt("db.port"),
				User:               viper.GetString("db.user"),
				Password:           viper.GetString("db.password"),
				DBName:             viper.GetString("db.dbname"),
				SSLmode:            viper.GetString("db.sslmode"),
				SetMaxIdleConns:    viper.GetInt("db.set_max_idle_conns"),
				SetMaxOpenConns:    viper.GetInt("db.set_max_open_conns"),
				SetConnMaxLifetime: viper.GetDuration("db.set_conn_max_lifetime"),
			},
			JWT: JWTConfig{
				SecretKey: viper.GetString("jwt.secret_key"),
				ExpiresIn: viper.GetDuration("jwt.expires_in"),
			},
			ADMIN: AdminConfig{
				UId: viper.GetString("admin.uid"),
			},
		}
	case "local":
		cfg = &config{
			API: APIConfig{
				Port: viper.GetString("local_api.port"),
			},
			DB: DBConfig{
				Host:               viper.GetString("local_db.host"),
				Port:               viper.GetInt("local_db.port"),
				User:               viper.GetString("local_db.user"),
				Password:           viper.GetString("local_db.password"),
				DBName:             viper.GetString("local_db.dbname"),
				SSLmode:            viper.GetString("local_db.sslmode"),
				SetMaxIdleConns:    viper.GetInt("local_db.set_max_idle_conns"),
				SetMaxOpenConns:    viper.GetInt("local_db.set_max_open_conns"),
				SetConnMaxLifetime: viper.GetDuration("local_db.set_conn_max_lifetime"),
			},
			JWT: JWTConfig{
				SecretKey: viper.GetString("local_jwt.secret_key"),
				ExpiresIn: viper.GetDuration("local_jwt.expires_in"),
			},
			ADMIN: AdminConfig{
				UId: viper.GetString("local_admin.uid"),
			},
		}
	case "debug":
		if os.Getenv("GINMODE") == "debug" {
			DBPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
			JWTExpiresIn, _ := strconv.Atoi(os.Getenv("JWT_EXPIRATION_TIME"))

			cfg.API.Port = os.Getenv("API_PORT")
			cfg.DB.Host = os.Getenv("DB_HOST")
			cfg.DB.Port = DBPort
			cfg.DB.User = os.Getenv("DB_USER")
			cfg.DB.Password = os.Getenv("DB_PASSWORD")
			cfg.DB.DBName = os.Getenv("DB_NAME")
			cfg.DB.SSLmode = "disable"
			cfg.DB.SetMaxIdleConns = 10
			cfg.DB.SetMaxOpenConns = 100
			cfg.DB.SetConnMaxLifetime = 60
			cfg.JWT.SecretKey = os.Getenv("JWT_SECRET_KEY")
			cfg.JWT.ExpiresIn = time.Duration(JWTExpiresIn)
			cfg.ADMIN.UId = os.Getenv("ADMIN_UID")

		}
	}
}
func GetConfig() *config {
	return cfg
}
