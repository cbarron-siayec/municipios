package insertTwitterPerformance

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertTwitterPerformance(friends int, favorites int, followers int, lists int, tweets int) {
	connString := "root:D3m0S14y3c@(127.0.0.1:3306)/municipios"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("CONNECTED")
	defer db.Close()
	stmt, err := db.Prepare("INSERT twitter_performance SET friends=?,favorites=?,followers=?,lists=?,tweets=?")
	if err != nil {
		panic(err.Error())
	}

	res, err := stmt.Exec(friends, favorites, followers, lists, tweets)
	if err != nil {
		panic(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(id)
}
