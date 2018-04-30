package twitterData

import(
  "fmt"
  "github.com/dghubble/go-twitter/twitter"
  "github.com/dghubble/oauth1"
)

type User struct {
	ContributorsEnabled            bool          `json:"contributors_enabled"`
	CreatedAt                      string        `json:"created_at"`
	DefaultProfile                 bool          `json:"default_profile"`
	DefaultProfileImage            bool          `json:"default_profile_image"`
	Description                    string        `json:"description"`
	Email                          string        `json:"email"`
	Entities                       *UserEntities `json:"entities"`
	FavouritesCount                int           `json:"favourites_count"`
	FollowRequestSent              bool          `json:"follow_request_sent"`
	Following                      bool          `json:"following"`
	FollowersCount                 int           `json:"followers_count"`
	FriendsCount                   int           `json:"friends_count"`
	GeoEnabled                     bool          `json:"geo_enabled"`
	ID                             int64         `json:"id"`
	IDStr                          string        `json:"id_str"`
	IsTranslator                   bool          `json:"is_translator"`
	Lang                           string        `json:"lang"`
	ListedCount                    int           `json:"listed_count"`
	Location                       string        `json:"location"`
	Name                           string        `json:"name"`
	Notifications                  bool          `json:"notifications"`
	ProfileBackgroundColor         string        `json:"profile_background_color"`
	ProfileBackgroundImageURL      string        `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHttps string        `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool          `json:"profile_background_tile"`
	ProfileBannerURL               string        `json:"profile_banner_url"`
	ProfileImageURL                string        `json:"profile_image_url"`
	ProfileImageURLHttps           string        `json:"profile_image_url_https"`
	ProfileLinkColor               string        `json:"profile_link_color"`
	ProfileSidebarBorderColor      string        `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string        `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string        `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool          `json:"profile_use_background_image"`
	Protected                      bool          `json:"protected"`
	ScreenName                     string        `json:"screen_name"`
	ShowAllInlineMedia             bool          `json:"show_all_inline_media"`
	Status                         *Tweet        `json:"status"`
	StatusesCount                  int           `json:"statuses_count"`
	Timezone                       string        `json:"time_zone"`
	URL                            string        `json:"url"`
	UtcOffset                      int           `json:"utc_offset"`
	Verified                       bool          `json:"verified"`
	WithheldInCountries            []string      `json:"withheld_in_countries"`
	WithholdScope                  string        `json:"withheld_scope"`
}

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
