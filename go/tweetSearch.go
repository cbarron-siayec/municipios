package main

import (
	"fmt"

	"database/sql"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	_ "github.com/go-sql-driver/mysql"
)

type Interactions struct {
	ID_Tweet         string                  `json:"id_tweet"`
	AuthorScreenName string                  `json:"author_screen_name"`
	LikeCount        int                     `json:"like_count"`
	RetweetCount     int                     `json:"retweet_count"`
	TweetContent     string                  `json:"tweet_content"`
	ReplyToScreename string                  `json:"reply_to_screename"`
	ReplyToStatusId  int64                   `json:"reply_to_status_id"`
	Lang             string                  `json:"lang"`
	Date             string                  `json:"fecha"`
	Hashtags         []twitter.HashtagEntity `json:"hashtags"`
	UserMentions     []twitter.MentionEntity `json:"user_mentions"`
}

func main() {
	config := oauth1.NewConfig("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5")
	token := oauth1.NewToken("2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e")
	connString := "root:D3m0S14y3c@(127.0.0.1:3306)/municipios"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	searchTweets, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{Query: "%40islacozumelmx", Count: 100, ResultType: "recent"})
	interactions := []Interactions{}
	for i := 0; i < len(searchTweets.Statuses); i++ {
		n := Interactions{
			ID_Tweet:         searchTweets.Statuses[i].IDStr,
			AuthorScreenName: searchTweets.Statuses[i].User.ScreenName,
			LikeCount:        searchTweets.Statuses[i].FavoriteCount,
			RetweetCount:     searchTweets.Statuses[i].RetweetCount,
			TweetContent:     searchTweets.Statuses[i].Text,
			ReplyToScreename: searchTweets.Statuses[i].InReplyToScreenName,
			ReplyToStatusId:  searchTweets.Statuses[i].InReplyToStatusID,
			Lang:             searchTweets.Statuses[i].Lang,
			Date:             searchTweets.Statuses[i].CreatedAt,
			Hashtags:         searchTweets.Statuses[i].Entities.Hashtags,
			UserMentions:     searchTweets.Statuses[i].Entities.UserMentions,
		}
		interactions = append(interactions, n)

	}
	if err != nil {
		panic(err.Error())
	}
	if resp == nil {
		fmt.Println("NO RESPONSE!")
	}
	for z := 0; z < len(interactions); z++ {
		stmt, err := db.Prepare("INSERT interactions SET id_tweet=?,tweet_content=?,author_screen_name=?,like_count=?,retweet_count=?,reply_to_screename=?,reply_to_status_id=?,lang=?,created_at=?")
		if err != nil {
			panic(err.Error())
		}
		res, err := stmt.Exec(interactions[z].ID_Tweet, interactions[z].TweetContent, interactions[z].AuthorScreenName, interactions[z].LikeCount, interactions[z].RetweetCount, interactions[z].ReplyToScreename, interactions[z].ReplyToStatusId, interactions[z].Lang, interactions[z].Date)
		if err != nil {
			panic(err.Error())
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err.Error())
		}
		println("Interactions last ID")
		println(id)
		for w := 0; w < len(interactions[z].Hashtags); w++ {
			stmt, err := db.Prepare("INSERT hashtags SET id_tweet=?,hashtag=?")
			if err != nil {
				panic(err.Error())
			}
			res, err := stmt.Exec(interactions[z].ID_Tweet, interactions[z].Hashtags[w].Text)
			if err != nil {
				panic(err.Error())
			}
			id, err := res.LastInsertId()
			if err != nil {
				panic(err.Error())
			}
			println("Interactions last ID")
			println(id)
		}
		for y := 0; y < len(interactions[z].UserMentions); y++ {
			stmt, err := db.Prepare("INSERT user_mentions SET id_tweet=?,id_user=?,screen_name=?")
			if err != nil {
				panic(err.Error())
			}
			res, err := stmt.Exec(interactions[z].ID_Tweet, interactions[z].UserMentions[y].ID, interactions[z].UserMentions[y].ScreenName)
			if err != nil {
				panic(err.Error())
			}
			id, err := res.LastInsertId()
			if err != nil {
				panic(err.Error())
			}
			println("Interactions last ID")
			println(id)
		}

	}
}
