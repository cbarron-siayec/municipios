package main

import (
	"fmt"
	"libs/getAllCandidatesTwitterData"
)

func main() {
	test := getAllCandidatesTwitterData.GetAllCandidatesTwitterData()
	fmt.Println(test)
}
