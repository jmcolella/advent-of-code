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

type dir struct {
	Name     string
	Parent   string
	Path     string
	Children []string
	Size     int
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

func dirSize(dirs []*dir, start *dir) int {
	sizeRegex, err := regexp.Compile(`(\d+) .+`)
	if err != nil {
		panic(err)
	}

	size := 0

	for _, c := range start.Children {
		data := sizeRegex.FindStringSubmatch(c)
		i, err := strconv.Atoi(data[1])
		if err != nil {
			panic(err)
		}

		size += i
	}

	for _, d := range dirs {
		if d.Parent == start.Name && d.Name != start.Name {
			size += dirSize(dirs, d)
		}
	}

	return size
}

func partOne(input []string) (int, error) {
	dirs := []*dir{}
	// currPath := []string{}
	cwd := "/"

	for _, line := range input {
		parsed, err := parseLine(line)
		if err != nil {
			return 0, err
		}

		fmt.Printf("\nline %v\n\n", line)

		fmt.Printf("\ncwd %v\n\n", cwd)

		currDirIdx := slices.IndexFunc(dirs, func(d *dir) bool {
			return d.Name == cwd
		})

		if parsed.Type == "cd" {
			if parsed.Name == ".." {
				currDir := dirs[currDirIdx]

				cwd = currDir.Parent

				// cwd = currPath[len(currPath)-2]
				// currPath = currPath[0 : len(currPath)-1]
			} else {
				dirIdx := slices.IndexFunc(dirs, func(d *dir) bool { return d.Name == parsed.Name })

				if dirIdx < 0 {
					dirs = append(dirs, &dir{
						Name:   parsed.Name,
						Parent: cwd,
						// Path: strings.Join(currPath, ""),
					})
				}

				// currPath = append(currPath, fmt.Sprintf("%s/", parsed.Name))

				cwd = parsed.Name
			}
		} else if parsed.Type == "dir" {
			dirIdx := slices.IndexFunc(dirs, func(d *dir) bool { return d.Name == parsed.Name && d.Parent == cwd })

			if dirIdx < 0 {
				dirs = append(dirs, &dir{
					Name:   parsed.Name,
					Parent: cwd,
					// Path: strings.Join(currPath, ""),
				})
			}
		} else if parsed.Type == "file" {
			fmt.Printf("cwd %v\n", cwd)

			fmt.Printf("line %v\n", line)
			currDir := dirs[currDirIdx]

			currDir.Children = append(currDir.Children, parsed.Name)
		}
	}

	sizeToCut := 0
	for _, d := range dirs {
		// size := dirSize(dirs, d)

		// if size <= 100000 {
		// 	sizeToCut += size
		// }

		fmt.Printf("dir Name: %v, Parent: %v, Children: %v\n", d.Name, d.Parent, d.Children)
	}

	return sizeToCut, nil
}

func partTwo(input []string) (int, error) {
	return 0, nil
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
