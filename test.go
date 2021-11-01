package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	userServices "go_project/services"
)


func test_err()  {
	defer func() {
		func() {
			recover()
		}()
	}()
	panic("llll")
}
func main2() {
	config := dbConfig{DbType: "mysql", Host: "zhqn.com", Username: "demo", Password: "demo", Database: "demo", Port: 3306}
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Database)
	fmt.Println("db url", url)
	db, err := sql.Open(config.DbType, url)
	if err != nil {
		panic(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)
	u := userServices.User{Name: "小明", Age: 18, RoleId: 1}
	userServices.Insert(&u, db)
	userServices.Insert(&u, db)

	test_err()
}
