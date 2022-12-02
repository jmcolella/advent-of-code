package main

import (
	"fmt"
	"main/day2"
)

func main() {
	// res, err := day1.New().Run()
	res, err := day2.New().Run()
	if err != nil {
		fmt.Printf("ERROR %v", err)
		return
	}

	fmt.Println(res)
}
