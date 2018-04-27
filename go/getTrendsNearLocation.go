package main
import (
  "fmt"
  "github.com/ChimeraCoder/anaconda"
)

func main(){
  api := anaconda.NewTwitterApiWithCredentials("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5", "2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e")
  trends := api.GetTrendsByPlace(24550705,nil)
  for page := range trends {
		//Print the current page of followers
		fmt.Println(page.Followers)
	}
}
