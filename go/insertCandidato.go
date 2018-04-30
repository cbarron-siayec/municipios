package insertCandidato

import (
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  _ "database/sql"
  "libs/sqlConnect"
)

func InsertCandidato(){
        sqlConnect.SqlConnect("root","D3m0S14y3c", "172.17.2.168",3306,"candidatos")
        stmt, err := db.Prepare("INSERT twitterData SET idCandidatos=?,friends=?,favorites=?,followers=?,lists=?,favorites=?,tweets=?")
        checkErr(err)

        res, err := stmt.Exec(1,1,1,1,1,1,1)
        checkErr(err)

        id, err := res.LastInsertId()
        checkErr(err)
        fmt.Printls(id)
}
