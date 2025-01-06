package database

import (
	"database/sql"
	"fmt"

	"github.com/herbetyp/crud-product-api/configs"
	_ "github.com/lib/pq"
)


func ConnectDB() *sql.DB {
	conf := configs.GetConfig()

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	// 	conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.DBName, conf.DB.SSLmode)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
	"localhost", 5432, "myuser", "mypassword", "postgres", "disable")
	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = conn.Ping()
	if err != nil {
		fmt.Printf("Error on connect database: %s", err.Error())
		panic(err)
	}
	
	fmt.Printf("Connected to database: : %s", conf.DB.DBName)
	return conn
}