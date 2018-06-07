package main

import (
	"libs/getTrendsByLocation"
	"libs/insertCandidato"
	"libs/insertTrends"
	"libs/insertTwitterPerformance"
	"libs/twitterData"
)

func main() {
	cozumel := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5", "2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e", "IslaCozumelMX", 0)
	amlo := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5", "2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e", "lopezobrador_", 1)
	rac := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5", "2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e", "RicardoAnayaC", 2)
	jamk := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5", "2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e", "JoseAMeadeK", 3)
	mz := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5", "2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e", "Mzavalagc", 4)
	bronco := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ", "gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5", "2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW", "Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e", "JaimeRdzNL", 5)
	trends := getTrendsByLocation.GetTrends(20.482623, -86.952691)
	// ORDEEN INSERT CANDIDATO: idCandidato int,friends int, favorites int, followers int, lists int, tweets int
	// NOMBRES ESTRUCTURA USER: IDCandidato,FriendsCount,FavouritesCount,FollowersCount,ListCount,NoTweets
	//PARA AMLO
	insertCandidato.InsertCandidato(amlo.IDCandidato, amlo.FriendsCount, amlo.FavouritesCount, amlo.FollowersCount, amlo.ListCount, amlo.NoTweets)
	//PARA RAC
	insertCandidato.InsertCandidato(rac.IDCandidato, rac.FriendsCount, rac.FavouritesCount, rac.FollowersCount, rac.ListCount, rac.NoTweets)
	//PARA JAMK
	insertCandidato.InsertCandidato(jamk.IDCandidato, jamk.FriendsCount, jamk.FavouritesCount, jamk.FollowersCount, jamk.ListCount, jamk.NoTweets)
	//PARA MZ
	insertCandidato.InsertCandidato(mz.IDCandidato, mz.FriendsCount, mz.FavouritesCount, mz.FollowersCount, mz.ListCount, mz.NoTweets)
	//PARA BRONCO
	insertCandidato.InsertCandidato(bronco.IDCandidato, bronco.FriendsCount, bronco.FavouritesCount, bronco.FollowersCount, bronco.ListCount, bronco.NoTweets)
	//PARA BRONCO
	insertTwitterPerformance.InsertTwitterPerformance(cozumel.FriendsCount, cozumel.FavouritesCount, cozumel.FollowersCount, cozumel.ListCount, cozumel.NoTweets)
	//Para Trends
	for i := 0; i < len(trends.Woeid); i++ {
		insertTrends.InsertTrends(trends.Name[i], trends.Volume[i], trends.Woeid[i])
	}
}
