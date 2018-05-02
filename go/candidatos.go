package main

import(
  _ "fmt"
  "net/http"
  "html/template"
   "libs/getAllCandidatesTwitterData"
)

func runIndex(w http.ResponseWriter, r *http.Request){
  test := getAllCandidatesTwitterData.GetAllCandidatesTwitterData()
  index, err := template.ParseFiles("static/graficos.html")
  index.Execute(w,test)
}

func main(){
  http.HandleFunc("/",runIndex)
  http.ListenAndServe(":80",nil)
}
