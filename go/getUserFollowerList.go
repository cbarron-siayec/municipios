package main

import (
	"database/sql"
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	_ "github.com/go-sql-driver/mysql"
)

type FollowersList struct {
	ID         string `json:"id"`
	ScreenName string `json:"screen_name"`
}

func main() {
	connString := "root:D3m0S14y3c@(172.17.2.154:3306)/municipios"
	db, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	config := oauth1.NewConfig("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5")
	token := oauth1.NewToken("2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e")
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	followers, resp, err := client.Followers.List(&twitter.FollowerListParams{UserID: 771547723707265025, Count: 200})
	followersList := []FollowersList{}
	for i := 0; i < len(followers.Users); i++ {
		n := FollowersList{ID: followers.Users[i].IDStr, ScreenName: followers.Users[i].ScreenName}
		followersList = append(followersList, n)
	}
	for z := 0; z < len(followersList)-1; z++ {
		stmt, err := db.Prepare("INSERT followers_list SET id_follower=?,screen_name=?")
		if err != nil {
			panic(err.Error())
		}

		res, err := stmt.Exec(followersList[z].ID, followersList[z].ScreenName)
		if err != nil {
			panic(err.Error())
		}
		id, err := res.LastInsertId()
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(id)
	}

	if err != nil {
		panic(err.Error())
	}
	if resp == nil {
		fmt.Println("NO RESPONSE!")
	}
}
