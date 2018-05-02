package main

import (
	"fmt"
	"libs/getAllCandidatesTwitterData"
)

func main() {
	test := getAllCandidatesTwitterData.GetAllCandidatesTwitterData()
	for i:= 0;i<len(test.IDTransaccion);i++{
		fmt.Println(test.IDTransaccion[i])
		fmt.Println(test.IDCandidato[i])
		fmt.Println(test.FriendsCount[i])
		fmt.Println(test.FavouritesCount[i])
		fmt.Println(test.FollowersCount[i])
		fmt.Println(test.ListCount[i])
		fmt.Println(test.NoTweets[i])
		fmt.Println(test.Date[i])
	}
}
