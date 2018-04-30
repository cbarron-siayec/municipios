package main

import (
	"fmt"
	"libs/twitterData"
	"libs/insertCandidato"
)

func main() {
	amlo := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ","gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5","2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW","Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e","lopezobrador_", 1)
	rac := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ","gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5","2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW","Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e","RicardoAnayaC", 2)
	jamk := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ","gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5","2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW","Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e","JoseAMeadeK", 3)
	mz := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ","gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5","2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW","Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e","Mzavalagc", 4)
	bronco := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ","gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5","2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW","Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e","JaimeRdzNL", 5)
	fmt.Println(amlo)
	fmt.Println(rac)
	fmt.Println(jamk)
	fmt.Println(mz)
	fmt.Println(bronco)
	insertCandidato.InsertCandidato(1,1,1,1,1,1)
}
