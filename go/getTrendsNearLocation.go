package main

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
 	"github.com/dghubble/oauth1"
)

func main(){
	config := oauth1.NewConfig(consumerKey,consumerSecret)
	token := oauth1.NewToken(accessToken,tokenSecret)
  	// http.Client will automatically authorize Requests
  	httpClient := config.Client(oauth1.NoContext, token)
  	// Twitter client
  	client := twitter.NewClient(httpClient)
	
	resp,err := client.Trends.Place(24550705,&twitter.TrendsPlaceParams{WOEID:24550705})
    	if err != nil {
      		panic(err.Error())
    	}
	fmt.Println(resp)		
}
