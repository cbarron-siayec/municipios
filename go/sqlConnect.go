package main

import (
	"fmt"
  "strings"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func sqlConnect(user string, password string, ip string, port string, database string) {
  connString := user+":"+password+"@("+ip+":"+port+"/"+database
  db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
}
func main() {
  sqlConnect(" 172.17.2.168","D3m0S14y3c","municipios")
}
