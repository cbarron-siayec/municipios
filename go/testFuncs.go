package main

import (
	"fmt"
	"libs/twitterData"
)

func main() {
	test := twitterData.GetData("D3hzZ2y3AlvG8o8Xw8SE2fxvZ","gpAkuWx7cyfGY6m6XffqcDyxNdwZRxhoCGDGS7ra54tCp1BHY5","2253585536-Yy1AxlyCNmEfEyJIOmhByqejD7XMXTpBIIHXoFW","Ha8IhwH4YElQwP2K3k3yKeT4EyDvZfd4EuxadGcddII9e","lopezobrador_", 1)
	fmt.Println(test)
}
