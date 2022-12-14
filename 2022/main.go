package main

import (
	"fmt"
	"main/day8"
)

func main() {
	// res, err := day1.New().Run()
	// res, err := day2.New().Run()
	// res, err := day3.New().Run()
	// res, err := day4.New().Run()
	// res, err := day5.New().Run()
	// res, err := day6.New().Run()
	// res, err := day7.New().Run()
	res, err := day8.New().Run()

	if err != nil {
		fmt.Printf("ERROR %v", err)
		return
	}

	fmt.Println(res)
}
