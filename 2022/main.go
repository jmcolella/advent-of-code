package main

import (
	"fmt"
	"main/day1"
)

func main() {
	res, err := day1.New().Run()
	if err != nil {
		fmt.Printf("ERROR %v", err)
		return
	}

	fmt.Println(res)
}
