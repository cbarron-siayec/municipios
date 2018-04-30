package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig("tyI3dExsT4Q5LjiA5Oxs9r2rL", "t6d19DqtSo7WjO8ChmX5dab8YOqw2kMv0yPHGW8EoAJYYGzL2i")
	token := oauth1.NewToken("119620082-ML9vIFx34BrmGNkKiK3scXwHjF7sisPfnwi1XqVs", "9jpCnF2ixF8dSlEaJrXyRRHbfhxNOwAekZveDtbAAH2Uw")
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	tweets, resp, err := client.Timelines.MentionTimeline(&twitter.MentionTimelineParams{Count: 200})
	for i := 0; i < len(tweets); i++ {
		fmt.Println(tweets[i].CreatedAtTime())
		fmt.Println(tweets[i].Text)
		fmt.Println(tweets[i].RetweetCount)
	}
	if err != nil {
		panic(err.Error())
	}
	if resp == nil {
		fmt.Println("NO RESPONSE!")
	}
}
