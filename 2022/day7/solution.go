package day7

import (
	"fmt"
	"main/readInput"
	"regexp"
	"strconv"

	"golang.org/x/exp/slices"
)

type Day7 struct{}

type linetype struct {
	Type string
	Name string
}

func parseLine(line string) (*linetype, error) {
	l := linetype{}
	commandRegex, err := regexp.Compile(`\$ (\w+)\s?(\D*)`)
	if err != nil {
		return &l, err
	}
	contentRegex, err := regexp.Compile(`(\w+) (\D+)`)
	if err != nil {
		return &l, err
	}

	if commandRegex.MatchString(line) {
		data := commandRegex.FindStringSubmatch(line)
		l.Type = data[1]
		if len(data) == 3 {
			l.Name = data[2]
		}
	} else if contentRegex.MatchString(line) {
		data := contentRegex.FindStringSubmatch(line)

		if data[1] == "dir" {
			l.Type = "dir"
			l.Name = data[2]
		} else {
			l.Type = "file"
			l.Name = line
		}
	}

	return &l, nil
}

func getSizes(parsedLines []linetype) []int {
	stack := []int{}
	sizes := []int{}

	for _, line := range parsedLines {
		if line.Type == "cd" {
			if line.Name == ".." {
				sizes = append(sizes, stack[len(stack)-1])
				stack = stack[0 : len(stack)-1]

				if len(stack) > 0 {
					stack[len(stack)-1] += sizes[len(sizes)-1]
				}
			} else {
				stack = append(stack, 0)
			}
		} else if line.Type == "file" {
			sizeRegex, err := regexp.Compile(`(\d+) .+`)
			if err != nil {
				panic(err)
			}
			data := sizeRegex.FindStringSubmatch(line.Name)
			i, err := strconv.Atoi(data[1])
			if err != nil {
				panic(err)
			}

			stack[len(stack)-1] += i
		}
	}

	for len(stack) > 0 {
		sizes = append(sizes, stack[len(stack)-1])
		stack = stack[0 : len(stack)-1]
		if len(stack) > 0 {
			stack[len(stack)-1] += sizes[len(sizes)-1]
		}
	}

	return sizes
}

func partOne(input []string) (int, error) {
	parsedLines := []linetype{}
	for _, line := range input {
		parsed, err := parseLine(line)
		if err != nil {
			return 0, err
		}
		parsedLines = append(parsedLines, *parsed)

	}

	sizes := getSizes(parsedLines)

	sizeToCut := 0
	for _, s := range sizes {
		if s <= 100000 {
			sizeToCut += s
		}
	}

	return sizeToCut, nil
}

func partTwo(input []string) (int, error) {
	parsedLines := []linetype{}
	for _, line := range input {
		parsed, err := parseLine(line)
		if err != nil {
			return 0, err
		}
		parsedLines = append(parsedLines, *parsed)

	}

	sizes := getSizes(parsedLines)

	slices.Sort(sizes)

	unused := 70000000 - sizes[len(sizes)-1]
	needed := 30000000 - unused

	potentialToCut := []int{}
	for _, s := range sizes {
		if s >= needed {
			potentialToCut = append(potentialToCut, s)
		}
	}

	slices.Sort(potentialToCut)

	return potentialToCut[0], nil
}

func (d *Day7) Run() (string, error) {
	input, err := readInput.New().ReadInputTxt("/day7/input.txt")
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

func New() *Day7 {
	return &Day7{}
}
