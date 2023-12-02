package aocutils

import (
	"bufio"
	"os"
)

// ParseInput takes a file path and returns all lines of the file and returns a slice of strings where each
// element is a line
func ParseInput(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func Sum(values []int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

func Product(values ...int) int {
	prod := 1
	for _, v := range values {
		prod *= v
	}
	return prod
}
