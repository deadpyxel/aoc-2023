package aocutils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func Split(input, sep string) (string, string, error) {
	parts := strings.Split(input, sep)
	if len(parts) != 2 {
		return "", "", fmt.Errorf("Split did not result in two parts, check separator and input string")
	}
	return parts[0], parts[1], nil
}

func StrToIntSlice(nums []string) ([]int, error) {
	var intNums []int
	for _, num := range nums {
		intNum, err := strconv.Atoi(num)
		if err != nil {
			return nil, err
		}
		intNums = append(intNums, intNum)
	}
	return intNums, nil
}
