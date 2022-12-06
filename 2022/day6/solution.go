package day6

import (
	"fmt"
	"main/readInput"
	"strings"
)

type Day6 struct{}

func detectMarker(input string, start, end int) int {
	if end > len(input) {
		return end
	}

	dupe := false
	letters := map[string]int{}
	packet := strings.Split(input[start:end], "")

	for _, p := range packet {
		if _, value := letters[p]; value {
			dupe = true
			letters[p] += 1
		} else {
			letters[p] = 1
		}
	}

	if !dupe {
		return end
	}

	return detectMarker(input, start+1, end+1)
}

func partOne(input []string) (int, error) {
	marker := detectMarker(input[0], 0, 4)

	return marker, nil
}

func partTwo(input []string) (int, error) {
	marker := detectMarker(input[0], 0, 14)

	return marker, nil
}

func (d *Day6) Run() (string, error) {
	input, err := readInput.New().ReadInputTxt("/day6/input.txt")
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

func New() *Day6 {
	return &Day6{}
}
