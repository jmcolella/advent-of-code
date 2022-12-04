package day3

import (
	"fmt"
	"main/readInput"
	"strings"
)

type Day3 struct{}

type alphabet struct {
	lower string
	upper string
}

func formatAlphabets() alphabet {
	lower := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

	upper := []string{}

	for _, letter := range lower {
		upper = append(upper, strings.ToUpper(letter))
	}

	return alphabet{
		lower: strings.Join(lower, ""),
		upper: strings.Join(upper, ""),
	}
}

func partOne(input []string) (int, error) {
	dupes := []string{}

	for _, line := range input {
		half := len(line) / 2
		first := line[0:half]
		second := line[half:]

		for _, item := range strings.Split(first, "") {
			if strings.Index(second, item) >= 0 {
				dupes = append(dupes, item)
				break
			}
		}
	}

	priority := 0

	alphabets := formatAlphabets()

	for _, letter := range dupes {
		lowerIdx := strings.Index(alphabets.lower, letter)
		upperIdx := strings.Index(alphabets.upper, letter)

		if lowerIdx > 0 {
			priority += lowerIdx + 1
		} else {
			priority += upperIdx + 26 + 1
		}
	}

	return priority, nil
}

func partTwo(input []string) (int, error) {
	groups := [][]string{}

	counter := 0

	for counter < len(input) {
		groups = append(groups, []string{
			input[counter],
			input[counter+1],
			input[counter+2],
		})

		counter += 3
	}

	badges := []string{}

	for _, group := range groups {
		first := group[0]
		second := group[1]
		third := group[2]

		for _, item := range strings.Split(first, "") {
			inSecond := strings.Index(second, item) >= 0
			inThird := strings.Index(third, item) >= 0

			if inSecond && inThird {
				badges = append(badges, item)
				break
			}
		}
	}

	priority := 0

	alphabets := formatAlphabets()

	for _, letter := range badges {
		lowerIdx := strings.Index(alphabets.lower, letter)
		upperIdx := strings.Index(alphabets.upper, letter)

		if lowerIdx > 0 {
			priority += lowerIdx + 1
		} else {
			priority += upperIdx + 26 + 1
		}
	}

	return priority, nil
}

func (d *Day3) Run() (string, error) {
	input, err := readInput.New().ReadInputTxt("/day3/input.txt")
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

func New() *Day3 {
	return &Day3{}
}
