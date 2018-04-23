package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func sqlConnect(user string, password string, ip string, port string, database string) {
  connString := user+":"+password+"@("+ip+":"+port+"/"+database
  db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("CONNECTED")
	defer db.Close()
}
