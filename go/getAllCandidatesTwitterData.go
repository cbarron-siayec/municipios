package main

import (
	"database/sql"
	"fmt"
	_ "log"

	_ "github.com/go-sql-driver/mysql"
)

//Structure passed from database
type TwitterUser struct {
	FavouritesCount int    `json:"favourites_count"`
	FollowersCount  int    `json:"followers_count"`
	FriendsCount    int    `json:"friends_count"`
	ID              int64  `json:"id"`
	ScreenName      string `json:"screen_name"`
	NoTweets        int    `json:"tweet_count"`
	ListCount       int    `json:"list_count"`
	IDCandidato     int    `json:"id_candidato"`
	Date            string `json:"fecha"`
	IDTransaccion   int    `json:"id_transaccion"`
}

func GetAllCandidatesTwitterData() []TwitterUser {
	// Open up our database connection.
	connString := "root:D3m0S14y3c@(172.17.2.168:3306)/candidatos"
	db, err := sql.Open("mysql", connString)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Execute the query
	results, err := db.Query("SELECT * FROM twitterData")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// ORDEEN INSERT CANDIDATO: idCandidato int,friends int, favorites int, followers int, lists int, tweets int
	// NOMBRES ESTRUCTURA USER: IDCandidato,FriendsCount,FavouritesCount,FollowersCount,ListCount,NoTweets
	var resultsTwitterUser []TwitterUser
	var int I
	I = 0
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&resultsTwitterUser[I].IDTransaccion, &resultsTwitterUser[I].IDCandidato[I], &resultsTwitterUser[I].FriendsCount, &resultsTwitterUser[I].FavouritesCount, &resultsTwitterUser[I].FollowersCount, &resultsTwitterUser[I].ListCount, &resultsTwitterUser[I].NoTweets, &resultsTwitterUser[I].Date)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		i++
	}
	return resultsTwitterUser
}

func main() {
	results := GetAllCandidatesTwitterData()
	fmt.Println(results[0].IDCandidato)
	fmt.Println(results[0].FollowersCount)
	fmt.Println(results[0].Date)
}
