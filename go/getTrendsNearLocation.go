package main
import (
  "fmt"
  "github.com/ChimeraCoder/anaconda"
)

func main(){
  api := anaconda.NewTwitterApiWithCredentials("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5", "2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e")
  searchResult, _ := api.GetSearch("golang", nil)
  for _ , tweet := range searchResult.Statuses {
      fmt.Println(tweet.Text)
  }
}
