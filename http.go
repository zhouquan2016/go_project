package main

import (
	"database/sql"
	"encoding/json"
	userServices "go_project/services"
	"io"
	"net/http"
)
var db *sql.DB

func mains() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/userAdd", userAdd)
	db = createDb()
	defer closeDb(db)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}

}


func userAdd(writer http.ResponseWriter, request *http.Request) {
	var buff [4098]byte
	readInt, err := request.Body.Read(buff[:])
	if readInt == 0 {
		writeBody(writer, 200,"ok")
		return
	}
	var u userServices.User
	err = json.Unmarshal(buff[:readInt], &u)
	if err != nil {
		panic(err)
	}
	_, err = userServices.Insert(&u, db)
	if err != nil {
		panic(err)
	}
	writeBody(writer, 200,"ok")
}

func writeBody(writer http.ResponseWriter, statusCode int, body interface{}) {
	writer.WriteHeader(statusCode)
	by, err := json.Marshal(body)
	if err != nil {
		panic(err)
	}
	writer.Write(by)
}

func hello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
