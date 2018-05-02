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

func GetAllCandidatesTwitterData() TwitterUser {
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
	var resultsTwitterUser TwitterUser
	for results.Next() {
		// for each row, scan the result into our tag composite object
		err = results.Scan(&resultsTwitterUser.IDTransaccion, &resultsTwitterUser.IDCandidato, &resultsTwitterUser.FriendsCount, &resultsTwitterUser.FavouritesCount, &resultsTwitterUser.FollowersCount, &resultsTwitterUser.ListCount, &resultsTwitterUser.NoTweets, &resultsTwitterUser.Date)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		fmt.Println(resultsTwitterUser.FollowersCount)
	}
	return resultsTwitterUser
}

func main() {
	fmt.Println("OK")
	GetAllCandidatesTwitterData()
}
