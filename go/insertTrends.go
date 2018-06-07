package insertTrends

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertCandidato(name string, volume int64, woeid int64) {
	connString := "root:D3m0S14y3c@(127.0.0.1:3306)/municipios"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("CONNECTED")
	defer db.Close()
	stmt, err := db.Prepare("INSERT trends SET name=?,volume=?,woeid=?")
	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Exec(name, volume, woeid)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(id)
}
