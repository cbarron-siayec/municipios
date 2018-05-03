package main

import (
	"fmt"
	"html/template"
	"libs/getAllCandidatesTwitterData"
	"net/http"
)

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

func runIndex(w http.ResponseWriter, r *http.Request) {
	twitterInfoAmlo := getAllCandidatesTwitterData.GetAllCandidatesTwitterData(1)
	twitterInfoAnaya := getAllCandidatesTwitterData.GetAllCandidatesTwitterData(2)
	twitterInfoMeade := getAllCandidatesTwitterData.GetAllCandidatesTwitterData(3)
	twitterInfoMargarita := getAllCandidatesTwitterData.GetAllCandidatesTwitterData(4)
	twitterInfoBronco := getAllCandidatesTwitterData.GetAllCandidatesTwitterData(5)
	condensedRes := []getAllCandidatesTwitterData.TwitterUser{}
	condensedRes = append(condensedRes, twitterInfoAmlo)
	condensedRes = append(condensedRes, twitterInfoAnaya)
	condensedRes = append(condensedRes, twitterInfoMeade)
	condensedRes = append(condensedRes, twitterInfoMargarita)
	condensedRes = append(condensedRes, twitterInfoBronco)

	index, err := template.ParseFiles("../static/html/candidatos/graficos.html")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(index.Execute(w, condensedRes))
}

func main() {
	http.HandleFunc("/", runIndex)
	http.ListenAndServe(":80", nil)
}
