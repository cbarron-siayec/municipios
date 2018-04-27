package main
import (
  "fmt"
  "github.com/ChimeraCoder/anaconda"
)

func main(){
  anaconda.SetConsumerKey("D3hzZ2y3AlvG8o8Xw8SE2fxvZ")
  anaconda.SetConsumerSecret("gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5")
  api := anaconda.NewTwitterApi("2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW","Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e")
  trends, err := api.GetTrendsByPlace(24550705,nil)
  if err != nil {
	  panic(err.Error())
  }
  for page := range trends.Trends {
		//Print the current page of followers
		fmt.Println(page)
  }
}
