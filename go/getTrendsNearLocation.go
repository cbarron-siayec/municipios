package main
import (
  "fmt"
  "github.com/ChimeraCoder/anaconda"
)

func main(){
  api := anaconda.NewTwitterApi("tyI3dExsT4Q5LjiA5Oxs9r2rL","t6d19DqtSo7WjO8ChmX5dab8YOqw2kMv0yPHGW8EoAJYYGzL2i")
  trends, err := api.GetTrendsByPlace(24550705,nil)
  if err != nil {
	  panic(err.Error())
  }
  for page := range trends.Trends {
		//Print the current page of followers
		fmt.Println(page)
  }
}
