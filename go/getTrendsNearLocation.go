package main
import (
  "fmt"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
)

func main(){
  
  config := oauth1.NewConfig("tyI3dExsT4Q5LjiA5Oxs9r2rL","t6d19DqtSo7WjO8ChmX5dab8YOqw2kMv0yPHGW8EoAJYYGzL2i")
  token := oauth1.NewToken("119620082-ML9vIFx34BrmGNkKiK3scXwHjF7sisPfnwi1XqVs","9jpCnF2ixF8dSlEaJrXyRRHbfhxNOwAekZveDtbAAH2Uw")
  // http.Client will automatically authorize Requests
  httpClient := config.Client(oauth1.NoContext, token)
  // Twitter client
  client := twitter.NewClient(httpClient)
  
   trends,resp,err := client.Trends.Closest(&twitter.ClosestParams{Lat:19.396224,Long:-99.159049})
    if err != nil {
      panic(err.Error())
    }
    if resp != nil {
      fmt.Println("No Response")
    }
    fmt.Println(trends[0])
    fmt.Println(trends[1])
    fmt.Println(trends[2])
    fmt.Println(trends[3])
}
