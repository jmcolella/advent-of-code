package day1

import (
	"fmt"
	"main/readInput"
	"sort"
	"strconv"
)

type Day1 struct{}

func partOne(input []string) (int64, error) {
	elvesCalories := []int64{}
	calorieCounter := 0

	for idx, i := range input {
		if len(i) == 0 {
			elvesCalories = append(elvesCalories, int64(calorieCounter))
			calorieCounter = 0
			continue
		}

		intI, err := strconv.Atoi(i)
		if err != nil {
			return 0, err
		}

		calorieCounter += intI

		if idx == len(input)-1 {
			elvesCalories = append(elvesCalories, int64(calorieCounter))
			calorieCounter = 0
		}
	}

	maxCalories := elvesCalories[0]

	for _, cal := range elvesCalories {
		if maxCalories < cal {
			maxCalories = cal
		}
	}

	return maxCalories, nil
}

func partTwo(input []string) (int, error) {
	elvesCalories := []int{}
	calorieCounter := 0

	for idx, i := range input {
		if len(i) == 0 {
			elvesCalories = append(elvesCalories, calorieCounter)
			calorieCounter = 0
			continue
		}

		intI, err := strconv.Atoi(i)
		if err != nil {
			return 0, err
		}

		calorieCounter += intI

		if idx == len(input)-1 {
			elvesCalories = append(elvesCalories, calorieCounter)
			calorieCounter = 0
		}
	}

	topCals := elvesCalories

	sort.Sort(sort.Reverse(sort.IntSlice(topCals)))

	totalCals := 0

	for idx, cal := range topCals {
		if idx < 3 {
			totalCals += cal
		}
	}

	return totalCals, nil
}

func (d *Day1) Run() (string, error) {
	input, err := readInput.New().ReadInputTxt("/day1/input.txt")
	if err != nil {
		return "", err
	}

	partOneAns, err := partOne(input)
	if err != nil {
		return "", err
	}

	partTwoAns, err := partTwo(input)
	if err != nil {
		return "", err
	}

	answers := fmt.Sprintf("Part One: %v\nPart Two: %v\n", partOneAns, partTwoAns)

	return answers, nil
}

func New() *Day1 {
	return &Day1{}
}
