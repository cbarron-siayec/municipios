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
	searchTweets, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{Query: "%40islacozumelmx"})
	for i := 0; i < len(searchTweets.Statuses); i++ {
		fmt.Println("              ")
		fmt.Println("--------------")
		fmt.Println("ID: ")
		fmt.Println(searchTweets.Statuses[i].ID)
		fmt.Println("Extended Tweet: ")
		fmt.Println(searchTweets.Statuses[i].ExtendedTweet)
		fmt.Println("Full Text: ")
		fmt.Println(searchTweets.Statuses[i].FullText)
		fmt.Println("User: ")
		fmt.Println(searchTweets.Statuses[i].User.ScreenName)
		fmt.Println("Created at time: ")
		fmt.Println(searchTweets.Statuses[i].CreatedAt)
		fmt.Println("Text: ")
		fmt.Println(searchTweets.Statuses[i].Text)
		fmt.Println("Retweeted: ")
		fmt.Println(searchTweets.Statuses[i].Retweeted)
		fmt.Println("Retweeted Status (Favourite Count): ")
		fmt.Println(searchTweets.Statuses[i].RetweetedStatus)
		fmt.Println("Hashtags: ")
		fmt.Println(searchTweets.Statuses[i].Entities.Hashtags)
		fmt.Println("Current User Retweet ID: ")
		fmt.Println(searchTweets.Statuses[i].CurrentUserRetweet)
		fmt.Println("User Mentions: ")
		fmt.Println(searchTweets.Statuses[i].Entities.Hashtags[i].Text)
		fmt.Println("Favorite Count: ")
		fmt.Println(searchTweets.Statuses[i].FavoriteCount)
		fmt.Println("Retweet Count: ")
		fmt.Println(searchTweets.Statuses[i].RetweetCount)
		fmt.Println("InReplyToScreenName: ")
		fmt.Println(searchTweets.Statuses[i].InReplyToScreenName)
		fmt.Println("InReplyToStatusId: ")
		fmt.Println(searchTweets.Statuses[i].InReplyToStatusID)
	}
	if err != nil {
		panic(err.Error())
	}
	if resp == nil {
		fmt.Println("NO RESPONSE!")
	}
}
