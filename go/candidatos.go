package main

import (
	"fmt"
	"html/template"
	"libs/getAllCandidatesTwitterData"
	"net/http"
)

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
	fmt.Println(twitterInfoAmlo)

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
