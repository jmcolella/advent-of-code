package day4

import (
	"fmt"
	"main/readInput"
	"strconv"
	"strings"
)

type Day4 struct{}

func assignment(ids string) (string, error) {
	bounds := strings.Split(ids, "-")
	lower, err := strconv.ParseInt(bounds[0], 10, 64)
	if err != nil {
		return "", nil
	}
	upper, err := strconv.ParseInt(bounds[1], 10, 64)
	if err != nil {
		return "", nil
	}

	a := ""
	c := lower

	for c <= upper {
		cStrv := strconv.FormatInt(c, 10)

		a += fmt.Sprintf("-%s-", cStrv)

		c++
	}

	// a should look like: -3--4--5--6-
	return a, nil
}

func formatAssignments(pair []string) (string, string, error) {
	first, err := assignment(pair[0])
	if err != nil {
		return "", "", err
	}
	second, err := assignment(pair[1])
	if err != nil {
		return "", "", err
	}

	return first, second, nil
}

func formatPairs(input []string) [][]string {
	pairs := [][]string{}

	for _, line := range input {
		pair := strings.Split(line, ",")

		pairs = append(pairs, []string{
			pair[0],
			pair[1],
		})
	}

	return pairs
}

func partOne(input []string) (int, error) {
	pairs := formatPairs(input)

	overlaps := 0

	for _, pair := range pairs {
		first, second, err := formatAssignments(pair)
		if err != nil {
			return 0, err
		}

		firstInSecond := strings.Index(second, first)
		secondInFirst := strings.Index(first, second)

		if firstInSecond >= 0 || secondInFirst >= 0 {
			overlaps++
		}
	}

	return overlaps, nil
}

func partTwo(input []string) (int, error) {
	pairs := formatPairs(input)

	overlaps := 0

	for _, pair := range pairs {
		first, second, err := formatAssignments(pair)
		if err != nil {
			return 0, err
		}

		for _, item := range strings.Split(first, "--") {
			// Isolate the actual assignment ID, ex. 3
			sanitized := strings.Replace(item, "-", "", len(item))

			// Match on the formatted assignment ID, ex. -3-
			match := strings.Index(second, fmt.Sprintf("-%s-", sanitized))

			if match >= 0 {
				overlaps++
				break
			}
		}
	}

	return overlaps, nil
}

func (d *Day4) Run() (string, error) {
	input, err := readInput.New().ReadInputTxt("/day4/input.txt")
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

func New() *Day4 {
	return &Day4{}
}
