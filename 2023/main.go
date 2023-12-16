package main

import (
	"2023/day1"
	"fmt"
)

func main() {
	day1 := day1.Day1{}

	ans, err := day1.Run()
	if err != nil {
		fmt.Printf("ERR %v\n", err)
	}

	fmt.Printf(ans)
}
