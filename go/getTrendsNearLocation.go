package main
import (
  "fmt"
  "github.com/ChimeraCoder/anaconda"
)

func main(){
  anaconda.SetConsumerKey("tyI3dExsT4Q5LjiA5Oxs9r2rL")
  anaconda.SetConsumerSecret("t6d19DqtSo7WjO8ChmX5dab8YOqw2kMv0yPHGW8EoAJYYGzL2i")
  api := anaconda.NewTwitterApi("119620082-ML9vIFx34BrmGNkKiK3scXwHjF7sisPfnwi1XqVs","9jpCnF2ixF8dSlEaJrXyRRHbfhxNOwAekZveDtbAAH2Uw")
  u := url.Parse("https://api.twitter.com/1.1/trends/place")
  v := u.Values{}
  v.Set("exclude", "hastags")
  trends, err := api.GetTrendsByPlace(24550705,v)
  if err != nil {
  	panic(err)
  }
		fmt.Println(trends)
}
