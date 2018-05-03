package main

import (
	"fmt"
	"html/template"
	"libs/getAllCandidatesTwitterData"
	"net/http"
)

func runIndex(w http.ResponseWriter, r *http.Request) {
	twitterInfoAmlo := getAllCandidatesTwitterData.GetAllCandidatesTwitterData(1)
	index, err := template.ParseFiles("../static/html/candidatos/graficos.html")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(index.Execute(w, twitterInfo))
}

func main() {
	http.HandleFunc("/", runIndex)
	http.ListenAndServe(":80", nil)
}
