package main

import(
  "fmt"
  "net/http"
  "html/template"
   "libs/getAllCandidatesTwitterData"
)

func runIndex(w http.ResponseWriter, r *http.Request){
  twitterInfo := getAllCandidatesTwitterData.GetAllCandidatesTwitterData()
  index, err := template.ParseFiles("static/graficos.html")
  if err != nil{
    panic(err.Error())
  }
  fmt.Println(index.Execute(w,twitterInfo))
}

func main(){
  http.HandleFunc("/",runIndex)
  http.ListenAndServe(":80",nil)
}
