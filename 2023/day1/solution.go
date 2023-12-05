package day1

import (
	"2023/readInput"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Day1 struct{}

func partOne(input []string) (int64, error) {
	total := int64(0)

	for _, line := range input {
		arr := strings.Split(line, "")
		left := 0
		right := len(arr) - 1

		tens := int64(0)
		ones := int64(0)

		for tens == 0 || ones == 0 {
			if t, err := strconv.ParseInt(arr[left], 10, 64); err != nil {
				left += 1
				continue
			} else if tens == 0 {
				tens = t * 10
			}
			if o, err := strconv.ParseInt(arr[right], 10, 64); err != nil {
				right -= 1
				continue
			} else if ones == 0 {
				ones = o
			}
		}

		total += (tens + ones)
	}

	return total, nil
}

func numKey(numStr string) int64 {
	nums := map[string]int64{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	numKeys := reflect.ValueOf(nums).MapKeys()

	for _, n := range numKeys {
		if strings.Contains(numStr, n.String()) {
			return nums[n.String()]
		}
	}

	return 0
}

func partTwo(input []string) (int64, error) {
	// nums := map[string]int64{
	// 	"one":   1,
	// 	"two":   2,
	// 	"three": 3,
	// 	"four":  4,
	// 	"five":  5,
	// 	"six":   6,
	// 	"seven": 7,
	// 	"eight": 8,
	// 	"nine":  9,
	// }
	// numKeys := reflect.ValueOf(nums).MapKeys()
	total := int64(0)

	for _, line := range input {
		arr := strings.Split(line, "")
		left := 0
		right := len(arr) - 1

		leftNum := []string{}
		rightNum := []string{}
		tens := int64(0)
		ones := int64(0)

		for tens == 0 || ones == 0 {
			if t, err := strconv.ParseInt(arr[left], 10, 64); err != nil {
				leftNum = append(leftNum, arr[left])

				num := numKey(strings.Join(leftNum, ""))

				if num != 0 {
					tens = num * 10
				} else {
					left += 1
					continue
				}
			} else if tens == 0 {
				tens = t * 10
			}
			if o, err := strconv.ParseInt(arr[right], 10, 64); err != nil {
				rightNum = append([]string{arr[right]}, rightNum...)

				num := numKey(strings.Join(rightNum, ""))

				if num != 0 {
					ones = num
				} else {
					right -= 1
					continue
				}
			} else if ones == 0 {
				ones = o
			}
		}

		total += (tens + ones)
	}

	return total, nil
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
