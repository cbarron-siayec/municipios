package main

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func main() {
	config := oauth1.NewConfig("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5")
	token := oauth1.NewToken("2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e")
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	userTweets, response, errs := client.Timelines.UserTimeline(&twitter.UserTimelineParams{ScreenName: "realDonaldTrump", Count: 50})
	if errs != nil {
		panic(errs.Error())
	}
	if response == nil {
		fmt.Println("NO RESPONSE!")
	}
	for l := 0; l < len(userTweets); l++ {
		tweet, resp, err := client.Statuses.Show(userTweets[l].ID, &twitter.StatusShowParams{ID: userTweets[l].ID})
		if err != nil {
			panic(err.Error())
		}
		if resp == nil {
			fmt.Println("NO RESPONSE!")
		}
		fmt.Println(l + 1)
		fmt.Println(tweet.Text)
		fmt.Println("Favorite Count: ")
		fmt.Println(tweet.FavoriteCount)
		fmt.Println("Retweet Count: ")
		fmt.Println(tweet.RetweetCount)
	}

}
