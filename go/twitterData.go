package twitterData

import(
  "fmt"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
)

func GetNumFollowers(consumerKey string, consumerSecret string,accessToken string, tokenSecret string,screenName string) int{
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
  noFollowers := show.FollowersCount
  noFriends := show.FriendsCount
  createdDate := show.CreatedAt
  email := show.Email
  noFavs := show.FavouritesCount
  noFollowing := show.Following
  geoEnabled := show.GeoEnabled
  userId := show.ID
  noListed := show.ListedCount
  loc := Location
  statusCount := StatusesCount
  timeZone := Timezone 
  url := URL
  if resp == nil {
    fmt.Println("No Response")
  }
   fmt.Println(noFriends)
   fmt.Println(createdDate)
   fmt.Println(email)
   fmt.Println(noFavs)
   fmt.Println(noFollowing)
   fmt.Println(geoEnabled)
   fmt.Println(userId)
   fmt.Println(noListed)
   fmt.Println(loc)
   fmt.Println(statusCount)
   fmt.Println(timeZone)
   fmt.Println(url)
  return noFollowers
}
