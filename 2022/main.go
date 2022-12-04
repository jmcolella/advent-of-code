package main

import (
	"fmt"
	"main/day4"
)

func main() {
	// res, err := day1.New().Run()
	// res, err := day2.New().Run()
	// res, err := day3.New().Run()
	res, err := day4.New().Run()

	if err != nil {
		fmt.Printf("ERROR %v", err)
		return
	}

	fmt.Println(res)
}
