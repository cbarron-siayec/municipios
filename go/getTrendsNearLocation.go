package main
import (
  "fmt"
  "github.com/ChimeraCoder/anaconda"
)

func main(){
  api := anaconda.NewTwitterApi("119620082-ML9vIFx34BrmGNkKiK3scXwHjF7sisPfnwi1XqVs","9jpCnF2ixF8dSlEaJrXyRRHbfhxNOwAekZveDtbAAH2Uw")
  trends, err := api.GetTrendsByPlace(24550705, nil)
  if err != nil {
  	panic(err)
  }
  for page := range trends.Trends {
		fmt.Println(page)
  }
}
