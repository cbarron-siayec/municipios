package main

import(
  "fmt"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
)



func main(){
config := oauth1.NewConfig("U93DMt4ydkuj7KjtTQOH3A7DT","dhpyVkyl6pxRDIvGsTcvCAQ9uNP2WJ9iFIAD2FlYCDY29nat1c")
token := oauth1.NewToken("2253585536-EAQIZqyv9Np8V1fJPcthYlT5TPvpacO0MK209j2", "OCgcI2poX3Y9fSVM0i2D0tWP9PepXmSiNJBRUMqAnmZU0")
// http.Client will automatically authorize Requests
httpClient := config.Client(oauth1.NoContext, token)
// Twitter client
client := twitter.NewClient(httpClient)

followers, resp, err := client.Followers.List(&twitter.FollowerListParams{})
  if err != nil {
    panic(err.Error())
  }
  //https://developer.twitter.com/en/docs/accounts-and-users/follow-search-get-users/api-reference/get-followers-list
  followersNames := followers.users.name
  fmt.Println(followers)
  fmt.Println("")
  fmt.Println(followersName)
  fmt.Println("")
  fmt.Println(resp)
}
