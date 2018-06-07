package getTrendsByLocation

import (
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type TrendingTopic struct {
	Name   []string `json:"name"`
	Volume []int64  `json:"volume"`
	Woeid  []int64  `json:"place"`
}

func GetTrends(lat float64, long float64) TrendingTopic {
	config := oauth1.NewConfig("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5")
	token := oauth1.NewToken("2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e")
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)

	location, response, errClos := client.Trends.Closest(&twitter.ClosestParams{Lat: lat, Long: long})
	if errClos != nil {
		panic(errClos.Error())
	}
	if response == nil {
		fmt.Println("No Response!")
	}
	var closestWoeids []int64

	for q := 0; q < len(location); q++ {
		z := location[q].WOEID
		closestWoeids = append(closestWoeids, z)
	}
	place := location[0].WOEID
	trends, resp, err := client.Trends.Place(closestWoeids[0], &twitter.TrendsPlaceParams{WOEID: closestWoeids[0]})
	if err != nil {
		panic(err.Error())
	}
	if resp == nil {
		fmt.Println("No Response!")
	}
	var topics TrendingTopic
	for i := 0; i < len(trends); i++ {
		for j := 0; j < len(trends[i].Trends); j++ {
			names := trends[i].Trends[j].Name
			volumes := trends[i].Trends[j].TweetVolume
			places := place
			topics.Name = append(topics.Name, names)
			topics.Volume = append(topics.Volume, volumes)
			topics.Woeid = append(topics.Woeid, places)
		}
	}
	return topics
}
