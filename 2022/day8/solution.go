package day8

import (
	"fmt"
	"main/readInput"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type Day8 struct{}

func createGrid(input []string) [][]int {
	grid := [][]int{}

	for _, line := range input {
		row := []int{}

		for _, i := range strings.Split(line, "") {
			num, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}

			row = append(row, num)
		}

		grid = append(grid, row)
	}

	return grid
}

func vis(grid [][]int, i int, j int, dir string, anchor int) bool {
	if dir == "up" {
		up := i - 1
		if up < 0 {
			return anchor > grid[i][j]
		}

		if anchor > grid[up][j] {
			return vis(grid, up, j, "up", anchor)
		}
	} else if dir == "down" {
		down := i + 1
		if down == len(grid) {
			return anchor > grid[i][j]
		}

		if anchor > grid[down][j] {
			return vis(grid, down, j, "down", anchor)
		}
	} else if dir == "left" {
		left := j - 1
		if left < 0 {
			return anchor > grid[i][j]
		}

		if anchor > grid[i][left] {
			return vis(grid, i, left, "left", anchor)
		}
	} else if dir == "right" {
		right := j + 1
		if right == len(grid[i]) {
			return anchor > grid[i][j]
		}

		if anchor > grid[i][right] {
			return vis(grid, i, right, "right", anchor)
		}
	}

	return false
}

func partOne(input []string) (int, error) {
	grid := createGrid(input)

	visible := 0

	for i := range grid {
		if i == 0 {
			visible += len(grid[i])
			continue
		}
		if i == len(grid)-1 {
			visible += len(grid[i])
			continue
		}

		for j := range grid[i] {
			if j == 0 || j == len(grid[i])-1 {
				visible += 1
				continue
			}

			item := grid[i][j]

			visUp := vis(grid, i, j, "up", item)
			visDown := vis(grid, i, j, "down", item)
			visLeft := vis(grid, i, j, "left", item)
			visRight := vis(grid, i, j, "right", item)

			if visUp || visDown || visLeft || visRight {
				visible += 1
			}

		}
	}

	return visible, nil
}

func visScore(grid [][]int, i int, j int, dir string, anchor int) int {
	score := 1

	if dir == "up" {
		up := i - 1
		if up < 0 {
			return 0
		}

		if anchor > grid[up][j] {
			score += visScore(grid, up, j, "up", anchor)
		}
	} else if dir == "down" {
		down := i + 1
		if down == len(grid) {
			return 0
		}

		if anchor > grid[down][j] {
			score += visScore(grid, down, j, "down", anchor)
		}
	} else if dir == "left" {
		left := j - 1
		if left < 0 {
			return 0
		}

		if anchor > grid[i][left] {
			score += visScore(grid, i, left, "left", anchor)
		}
	} else if dir == "right" {
		right := j + 1
		if right == len(grid[i]) {
			return 0
		}

		if anchor > grid[i][right] {
			score += visScore(grid, i, right, "right", anchor)
		}
	}

	return score
}

func partTwo(input []string) (int, error) {
	grid := createGrid(input)

	scores := []int{}

	for i := range grid {
		for j := range grid[i] {
			item := grid[i][j]

			visUp := visScore(grid, i, j, "up", item)
			visDown := visScore(grid, i, j, "down", item)
			visLeft := visScore(grid, i, j, "left", item)
			visRight := visScore(grid, i, j, "right", item)

			score := visUp * visDown * visLeft * visRight

			scores = append(scores, score)
		}
	}

	slices.Sort(scores)

	return scores[len(scores)-1], nil
}

func (d *Day8) Run() (string, error) {
	input, err := readInput.New().ReadInputTxt("/day8/input.txt")
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

func New() *Day8 {
	return &Day8{}
}
