package main

import (
  "fmt"
  "github.com/go-sql-driver/mysql"
  "database/sql"
)

func main(){
        connString := "root:D3m0S14y3c@(172.17.2.168:3306/candidatos)"
        db, err := sql.Open("mysql", connString)
        if err != nil {
            panic(err.Error())
        }
        fmt.Println("CONNECTED")
        defer db.Close()
        stmt, err := db.Prepare("INSERT twitterData SET idCandidatos=?,friends=?,favorites=?,followers=?,lists=?,favorites=?,tweets=?")
        checkErr(err)

        res, err := stmt.Exec(1,1,1,1,1,1,1)
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)
        fmt.Printls(id)
}
