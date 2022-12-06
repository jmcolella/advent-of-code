package day5

import (
	"fmt"
	"main/readInput"
	"regexp"
	"strconv"
)

type Day5 struct{}

func formatInput(input []string) ([][]string, [][]int, error) {
	cargo := [][]string{}
	instructions := [][]int{}

	instructionsRegexp, err := regexp.Compile(`move (\d+) from (\d+) to (\d+)`)
	if err != nil {
		return cargo, instructions, err
	}

	for _, line := range input {
		m := instructionsRegexp.FindStringSubmatch(line)
		if len(m) > 0 {
			amount, err := strconv.Atoi(m[1])
			if err != nil {
				return cargo, instructions, err
			}
			start, err := strconv.Atoi(m[2])
			if err != nil {
				return cargo, instructions, err
			}
			finish, err := strconv.Atoi(m[3])
			if err != nil {
				return cargo, instructions, err
			}

			instructions = append(instructions, []int{amount, start, finish})
			continue
		}

		if len(line) == 0 {
			continue
		}

		c := 0

		temp := []string{}
		for c < len(line) {
			item := line[c : c+3]

			temp = append(temp, item)

			c += 4
		}

		cargo = append(cargo, temp)
	}

	cargo = cargo[0 : len(cargo)-1]

	stacks := [][]string{}

	i := 0

	itemRegexp, err := regexp.Compile(`[A-Z]`)
	if err != nil {
		return stacks, instructions, err
	}

	for i < len(cargo[0]) {
		temp := []string{}

		j := 0
		for j < len(cargo) {
			item := cargo[j][i]
			if itemRegexp.MatchString(item) {
				temp = append(temp, item)
			}

			j++
		}

		stacks = append(stacks, temp)
		i++
	}

	return stacks, instructions, nil
}

func makeCode(stacks [][]string) (string, error) {
	code := ""

	crateRegexp, err := regexp.Compile(`\[(\w)\]`)
	if err != nil {
		return "", err
	}
	for _, s := range stacks {
		m := crateRegexp.FindStringSubmatch(s[0])

		code += m[1]
	}

	return code, nil
}

func partOne(input []string) (string, error) {
	stacks, instructions, err := formatInput(input)
	if err != nil {
		return "", err
	}

	for _, i := range instructions {
		amount := i[0]
		start := i[1] - 1
		finish := i[2] - 1

		items := stacks[start][0:amount]
		stacks[start] = stacks[start][amount:]

		for _, i := range items {
			stacks[finish] = append([]string{i}, stacks[finish]...)
		}
	}

	code, err := makeCode(stacks)
	if err != nil {
		return "", err
	}

	return code, nil
}

func partTwo(input []string) (string, error) {
	stacks, instructions, err := formatInput(input)
	if err != nil {
		return "", err
	}

	for _, i := range instructions {
		amount := i[0]
		start := i[1] - 1
		finish := i[2] - 1

		items := stacks[start][0:amount]
		stacks[start] = stacks[start][amount:]

		c := len(items) - 1

		for c >= 0 {
			stacks[finish] = append([]string{items[c]}, stacks[finish]...)
			c--
		}
	}

	code, err := makeCode(stacks)
	if err != nil {
		return "", err
	}

	return code, nil
}

func (d *Day5) Run() (string, error) {
	input, err := readInput.New().ReadInputTxt("/day5/input.txt")
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

func New() *Day5 {
	return &Day5{}
}
