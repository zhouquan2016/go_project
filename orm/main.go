package orm

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
)

var db *sql.DB
var configPath string
var config *DbConfig

func init()  {
	configPath = "db.json"
}
func getConf(path string)  *DbConfig{
	config, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var dbConfig DbConfig
	json.Unmarshal(config, &dbConfig)
	return &dbConfig
}

func registerConfigPath(path string)  {
	configPath = path
}
func InitOrm()  {
	config = getConf(configPath)
	//sql.Open()
}

