package sqlConnect

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func SqlConnect(user string, password string, ip string, port string, database string) {
  connString := user+":"+password+"@("+ip+":"+port+"/"+database
  DB, ERR := sql.Open("mysql", connString)
	if ERR != nil {
		panic(ERR.Error())
	}
	fmt.Println("CONNECTED")
	defer DB.Close()
}
