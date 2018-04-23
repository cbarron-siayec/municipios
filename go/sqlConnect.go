package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func sqlConnect(user string, password string, ip string, port string, database string) {
  connString := user+":"+password+"@("+ip+":"+port+"/"+database
  fmt.Println(connString)
  db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("CONNECTED")
	defer db.Close()
}
func main() {
  sqlConnect("root","D3m0S14y3c","172.17.2.168","3306","municipios")
}
