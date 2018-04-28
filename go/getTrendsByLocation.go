package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
 	"github.com/dghubble/oauth1"
)

func main(){
	config := oauth1.NewConfig("D3hzZ2y3AlvG8o8Xw8SE2fxvZ","gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5")
	token := oauth1.NewToken("2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW","Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e")
  	// http.Client will automatically authorize Requests
  	httpClient := config.Client(oauth1.NoContext, token)
  	// Twitter client
  	client := twitter.NewClient(httpClient)
	trends,resp,err := client.Trends.Place(116545,&twitter.TrendsPlaceParams{WOEID:116545})
	
    	if err != nil {
      		panic(err.Error())
    	}
	if resp != nil {
		fmt.Println(resp)
    	}
	for i:=0;i<=len(trends);i++ {
		fmt.Println(trends[i])
	}
}
