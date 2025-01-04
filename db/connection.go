package db

import (
	"database/sql"
	"fmt"

	"github.com/herbetyp/crud-product-api/configs"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	conf := configs.GetConfig()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.DBName, conf.DB.SSLmode)
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to " + conf.DB.DBName)
	return conn, nil
}
