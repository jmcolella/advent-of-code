package readInput

import (
	"bufio"
	"encoding/csv"
	"os"
)

type ReadInput struct{}

func (ri *ReadInput) ReadInputTxt(filePath string) ([]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	txtFile, err := os.Open(wd + filePath)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(txtFile)

	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func (ri *ReadInput) ReadInputCsv(filePath string) ([][]string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	csvFile, err := os.Open(wd + filePath)
	if err != nil {
		return nil, err
	}

	csvReader := csv.NewReader(csvFile)
	return csvReader.ReadAll()
}

func New() *ReadInput {
	return &ReadInput{}
}
