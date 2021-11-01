package main

import (
	"database/sql"
	"fmt"
)

type dbConfig struct {
	Host     string
	Username string
	Password string
	Database string
	Port   int16
	DbType string
}

func createDb() *sql.DB {

	config := dbConfig{DbType: "mysql", Host: "zhqn.com", Username: "demo", Password: "demo", Database: "demo", Port: 3306}
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	fmt.Println("db url", url)
	db, err := sql.Open(config.DbType, url)
	if err != nil {
		panic(err)
	}
	return db
}

func closeDb(db *sql.DB) {
	db.Close()
}