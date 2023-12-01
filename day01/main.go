package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
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

func ProccesLines(lines []string) ([]int, error) {
	var numbers []int
	re := regexp.MustCompile("[0-9]")
	for _, line := range lines {
		nums := re.FindAllString(line, -1)

		if len(nums) == 0 {
			continue
		}

		first := nums[0]
		last := nums[len(nums)-1]

		number, err := strconv.Atoi(first + last)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}
	return numbers, nil
}

func solverPart1(lines []string) (int, error) {

	numbers, err := ProccesLines(lines)
	if err != nil {
		return 0, err
	}
	var sol int
	for _, n := range numbers {
		sol += n
	}
	return sol, nil
}

func main() {
	filepath := "day01.txt"
	lines, err := ParseInput(filepath)
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}

	sol01, err := solverPart1(lines)
	if err != nil {
		log.Fatalf("Error while solving problem 1-1: %v", err)
	}
	fmt.Printf("Solution for 1-1: %d\n", sol01)

}
