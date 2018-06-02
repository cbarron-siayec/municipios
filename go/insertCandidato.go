package insertCandidato

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "database/sql"
)

func InsertCandidato(idCandidato int,friends int, favorites int, followers int, lists int, tweets int){
        connString := "root:D3m0S14y3c@(127.0.0.1:3306)/candidatos"
        db, err := sql.Open("mysql", connString)
        if err != nil {
            panic(err.Error())
        }
        fmt.Println("CONNECTED")
        defer db.Close()
        stmt, err := db.Prepare("INSERT twitterData SET idCandidatos=?,friends=?,favorites=?,followers=?,lists=?,tweets=?")
        if err != nil {
            panic(err.Error())
        }

        res, err := stmt.Exec(idCandidato,friends,favorites,followers,lists,tweets)
        if err != nil {
            panic(err.Error())
        }

        id, err := res.LastInsertId()
        if err != nil {
            panic(err.Error())
        }
        fmt.Println(id)
}
