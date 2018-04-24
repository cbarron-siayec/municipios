package getNumFollowers

import(
  "fmt"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
)



func HowManyFollowers(conumerKey string, consumerSecret string,accessToken string, tokenSecret string){
config := oauth1.NewConfig(consumerKey,consumerSecret)
token := oauth1.NewToken(accessToken,tokenSecret)
// http.Client will automatically authorize Requests
httpClient := config.Client(oauth1.NoContext, token)
// Twitter client
client := twitter.NewClient(httpClient)

followers, resp, err := client.Followers.List(&twitter.FollowerListParams{})
  if err != nil {
    panic(err.Error())
  }
  //https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-list
  noFollowers := len(followers.Users)
  if resp == nil {
    fmt.Println("No Response")
  }
  return noFollowers

}
