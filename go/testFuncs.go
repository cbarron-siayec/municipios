package main

import (
	"libs/getAllCandidatesTwitterData"
)

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

func main() {
	test := getAllCandidatesTwitterData.GetAllCandidatesTwitterData()
}
