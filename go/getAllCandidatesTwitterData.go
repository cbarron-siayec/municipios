package getAllCandidatesTwitterData

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Structure passed from database
type TwitterUser struct {
	FavouritesCount []float64 `json:"favourites_count"`
	FollowersCount  []float64 `json:"followers_count"`
	FriendsCount    []float64 `json:"friends_count"`
	ID              []float64 `json:"id"`
	ScreenName      []string  `json:"screen_name"`
	NoTweets        []float64 `json:"tweet_count"`
	ListCount       []float64 `json:"list_count"`
	IDCandidato     []int     `json:"id_candidato"`
	Date            []string  `json:"fecha"`
	IDTransaccion   []int     `json:"id_transaccion"`
}

func GetAllCandidatesTwitterData(idCandidato int) TwitterUser {
	// Open up our database connection.
	connString := "root:D3m0S14y3c@(127.0.0.1:3306)/candidatos"
	db, err := sql.Open("mysql", connString)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Execute the query
	results, err := db.Query("SELECT * FROM twitterData WHERE idCandidatos=?", idCandidato)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// ORDEEN INSERT CANDIDATO: idCandidato int,friends int, favorites int, followers int, lists int, tweets int
	// NOMBRES ESTRUCTURA USER: IDCandidato,FriendsCount,FavouritesCount,FollowersCount,ListCount,NoTweets
	var resultsTwitterUser TwitterUser
	for results.Next() {
		var idtransaccion, idcandidatos int
		var friends, followers, favorites, lists, tweets float64
		var fecha string
		// for each row, scan the result into our tag composite object
		err = results.Scan(&idtransaccion, &idcandidatos, &friends, &favorites, &followers, &lists, &tweets, &fecha)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		resultsTwitterUser.IDTransaccion = append(resultsTwitterUser.IDTransaccion, idtransaccion)
		resultsTwitterUser.IDCandidato = append(resultsTwitterUser.IDCandidato, idcandidatos)
		resultsTwitterUser.FriendsCount = append(resultsTwitterUser.FriendsCount, friends)
		resultsTwitterUser.FavouritesCount = append(resultsTwitterUser.FavouritesCount, favorites)
		resultsTwitterUser.FollowersCount = append(resultsTwitterUser.FollowersCount, followers)
		resultsTwitterUser.ListCount = append(resultsTwitterUser.ListCount, lists)
		resultsTwitterUser.NoTweets = append(resultsTwitterUser.NoTweets, tweets)
		resultsTwitterUser.Date = append(resultsTwitterUser.Date, fecha)
	}
	return resultsTwitterUser
}
