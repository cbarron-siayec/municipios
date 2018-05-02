package main

import(
  "fmt"
  "net/http"
  "html/tmeplate"
   "libs/getAllCandidatesTwitterData"
)

func runIndex(w http.ResponseWriter, r *http.Request){
  test := getAllCandidatesTwitterData.GetAllCandidatesTwitterData()
  index, err := template.ParseFiles("static/graficos.html")
  index.execute(w,test)
}

func main(){
  http.HandleFunc("/",runIndex)
  http.ListenAndServer(:80,nil)
}
