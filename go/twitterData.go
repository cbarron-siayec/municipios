package twitterData

import(
  "fmt"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
)

type User struct {
	FavouritesCount                int           `json:"favourites_count"`
	FollowersCount                 int           `json:"followers_count"`
	FriendsCount                   int           `json:"friends_count"`
	ID                             int64         `json:"id"`
	ScreenName                     string        `json:"screen_name"`
	NoTweets                  int           `json:"statuses_count"`
	ListCount		       int	     'json:"list_count"'
	IDCandidato		       int	     'json:"IDCandidato"'
}

func GetData(consumerKey string, consumerSecret string,accessToken string, tokenSecret string,screenName string, idCandidato int) User{
  config := oauth1.NewConfig(consumerKey,consumerSecret)
  token := oauth1.NewToken(accessToken,tokenSecret)
  // http.Client will automatically authorize Requests
  httpClient := config.Client(oauth1.NoContext, token)
  // Twitter client
  client := twitter.NewClient(httpClient)
  
   show,resp,err := client.Users.Show(&twitter.UserShowParams{ScreenName:screenName})
    if err != nil {
      panic(err.Error())
    }
  
  //https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-list 
  twitterInfo := User{FavouritesCount: show.FavouritesCount, FollowersCount: show.FollowersCount,FriendsCount:show.FriendsCount,ScreenName:show.ScreenName,ID:show.ID,ListCount:show.ListedCount,noTweets:show.StatusesCount,IDCandidato:idCandidato}
  if resp == nil {
    fmt.Println("No Response")
  }
  return twitterInfo
}
