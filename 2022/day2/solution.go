package day2

import (
	"fmt"
	"main/readInput"
	"strconv"
	"strings"
)

type Day2 struct{}

func createRules() map[string]map[string]string {
	rules := make(map[string]map[string]string)
	rules["X"] = map[string]string{}
	rules["Y"] = map[string]string{}
	rules["Z"] = map[string]string{}

	rules["X"]["A"] = "tie"
	rules["X"]["B"] = "lose"
	rules["X"]["C"] = "win"
	rules["X"]["Points"] = "1"

	rules["Y"]["A"] = "win"
	rules["Y"]["B"] = "tie"
	rules["Y"]["C"] = "lose"
	rules["Y"]["Points"] = "2"

	rules["Z"]["A"] = "lose"
	rules["Z"]["B"] = "win"
	rules["Z"]["C"] = "tie"
	rules["Z"]["Points"] = "3"

	return rules
}

func partOne(input [][]string) (int, error) {
	rules := createRules()

	scores := []int{}

	for _, game := range input {
		gameTuple := strings.Split(game[0], " ")
		opponent := gameTuple[0]
		choice := gameTuple[1]

		outcome := rules[choice][opponent]
		points, err := strconv.Atoi(rules[choice]["Points"])
		if err != nil {
			return 0, nil
		}

		if outcome == "win" {
			score := 6 + points
			scores = append(scores, score)
		} else if outcome == "lose" {
			score := points
			scores = append(scores, score)
		} else {
			score := 3 + points
			scores = append(scores, score)
		}
	}

	total := 0

	for _, s := range scores {
		total += s
	}

	return total, nil
}

func createRulesTwo() map[string]map[string]string {
	rules := make(map[string]map[string]string)
	rules["A"] = map[string]string{}
	rules["B"] = map[string]string{}
	rules["C"] = map[string]string{}

	rules["A"]["tie"] = "A"
	rules["A"]["lose"] = "B"
	rules["A"]["win"] = "C"
	rules["A"]["Points"] = "1"

	rules["B"]["win"] = "A"
	rules["B"]["tie"] = "B"
	rules["B"]["lose"] = "C"
	rules["B"]["Points"] = "2"

	rules["C"]["lose"] = "A"
	rules["C"]["win"] = "B"
	rules["C"]["tie"] = "C"
	rules["C"]["Points"] = "3"

	return rules
}

func partTwo(input [][]string) (int, error) {
	rules := createRulesTwo()

	scores := []int{}

	for _, game := range input {
		gameTuple := strings.Split(game[0], " ")
		opponent := gameTuple[0]
		outcome := gameTuple[1]

		// Opponent wins
		if outcome == "X" {
			choice := rules[opponent]["win"]
			points, err := strconv.Atoi(rules[choice]["Points"])
			if err != nil {
				return 0, nil
			}

			score := points
			scores = append(scores, score)

			// Opponent ties
		} else if outcome == "Y" {
			choice := rules[opponent]["tie"]
			points, err := strconv.Atoi(rules[choice]["Points"])
			if err != nil {
				return 0, nil
			}

			score := 3 + points
			scores = append(scores, score)

			// Opponent loses
		} else {
			choice := rules[opponent]["lose"]
			points, err := strconv.Atoi(rules[choice]["Points"])
			if err != nil {
				return 0, nil
			}

			score := 6 + points
			scores = append(scores, score)
		}
	}

	total := 0

	for _, s := range scores {
		total += s
	}

	return total, nil
}

func (d *Day2) Run() (string, error) {
	input, err := readInput.New().ReadInputCsv("/day2/input.csv")
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

func New() *Day2 {
	return &Day2{}
}
