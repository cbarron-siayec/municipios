package getFollowersTest

import(
   "fmt"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
)

func getNumFollowers(consumerKey string, consumerSecret string,accessToken string, tokenSecret string,screenName string) int{
  config := oauth1.NewConfig(consumerKey,consumerSecret)
  token := oauth1.NewToken(accessToken,tokenSecret)
  // http.Client will automatically authorize Requests
  httpClient := config.Client(oauth1.NoContext, token)
  // Twitter client
  client := twitter.NewClient(httpClient)
  
   show,resp,err := client.Users.Show(&twitter.UserShowParams{ScreenName:"realDonaldTrump"})
    if err != nil {
      panic(err.Error())
    }
  
  //https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-list
  noFollowers := len(show.FollowersCount)
  if resp == nil {
    fmt.Println("No Response")
  }
  return noFollowers
}
