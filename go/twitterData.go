package twitterData

import(
  "fmt"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
)

func GetData(consumerKey string, consumerSecret string,accessToken string, tokenSecret string,screenName string) User{
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
  //noFollowers := show.FollowersCount
  //noFriends := show.FriendsCount
  //createdDate := show.CreatedAt
  //email := show.Email
  //noFavs := show.FavouritesCount
  //noFollowing := show.Following
  //geoEnabled := show.GeoEnabled
  //userId := show.ID
  //noListed := show.ListedCount
  //noTweets := show.StatusesCount
  if resp == nil {
    fmt.Println("No Response")
  }
  return show  
}
