package main

import (
	"fmt"
	"strconv"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

func GetTrends(WOEID int64) [50]string {
	config := oauth1.NewConfig("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5")
	token := oauth1.NewToken("2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e")
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)
	// Twitter client
	client := twitter.NewClient(httpClient)
	trends, resp, err := client.Trends.Place(WOEID, &twitter.TrendsPlaceParams{WOEID: WOEID})

	if err != nil {
		panic(err.Error())
	}
	if resp == nil {
		fmt.Println("No Response!")
	}
	var top50 [50]string
	for i := 0; i < len(trends); i++ {
		for j := 0; j < len(trends[i].Trends); j++ {
			names := trends[i].Trends[j].Name
			volumes := strconv.FormatInt(trends[i].Trends[j].TweetVolume, 10)
			concRes := "Name: " + names + ", Volume: " + volumes
			top50[j] = concRes
		}
	}
	return top50
}
func main() {
	resArray := GetTrends(151582)
	for k := 0; k < len(resArray); k++ {
		if resArray[k] != "" {
			fmt.Println(resArray[k])
		}
	}
}
