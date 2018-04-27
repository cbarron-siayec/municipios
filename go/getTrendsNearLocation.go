package main
import (
  "fmt"
  "github.com/ChimeraCoder/anaconda"
)

func main(){
  api := anaconda.NewTwitterApiWithCredentials("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5", "2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e")
  api, trends, resp := anaconda.GetTrendsByPlace(24550705,"")
  if resp != nil {
    fmt.Println("Error")
  }
  fmt.Println(trends)
  fmt.Println(resp)
}
