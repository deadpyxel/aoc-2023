package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var strToDigit map[string]string

func createMap() {
	// Custom map for string to digit
	strToDigit = make(map[string]string)
	strToDigit["one"] = "1"
	strToDigit["two"] = "2"
	strToDigit["three"] = "3"
	strToDigit["four"] = "4"
	strToDigit["five"] = "5"
	strToDigit["six"] = "6"
	strToDigit["seven"] = "7"
	strToDigit["eight"] = "8"
	strToDigit["nine"] = "9"
	strToDigit[""] = "0"

}

func extract(input string, re *regexp.Regexp) (string, string) {
	var first, last string
	for i := range input {
		matches := re.FindString(input[:i])
		if len(matches) > 0 {
			first = matches
			break
		}
	}
	for i := len(input) - 1; i >= 0; i-- {
		curr := input[i:]
		matches := re.FindString(curr)
		if len(matches) > 0 {
			last = matches
			break
		}
	}

	return first, last
}

func asInteger(first, last string) (int, error) {

	if len(first) > 1 || len(first) == 0 {
		first = strToDigit[first]
	}
	if len(last) > 1 || len(last) == 0 {
		last = strToDigit[last]
	}

	number, err := strconv.Atoi(first + last)
	if err != nil {
		return 0, err
	}
	return number, nil
}

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

func processLine(line string) string {
	subMap := map[string]string{
		"one":   "o1e",
		"two":   "t2o",
		"three": "t3e",
		"four":  "f4r",
		"five":  "f5e",
		"six":   "s6x",
		"seven": "s7n",
		"eight": "e8t",
		"nine":  "n9e",
	}
	for spelled, repr := range subMap {
		line = strings.ReplaceAll(line, spelled, repr)
	}
	return line
}

func ProcessLines(lines []string, regex string) ([]int, error) {
	var numbers []int
	re := regexp.MustCompile(regex)
	for _, line := range lines {
		first, last := extract(line, re)

		number, err := asInteger(first, last)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}
	return numbers, nil
}

func solvePart1(lines []string) (int, error) {

	numbers, err := ProcessLines(lines, "[0-9]")
	if err != nil {
		return 0, err
	}
	var sol int
	for _, n := range numbers {
		sol += n
	}
	return sol, nil
}

func solvePart2(lines []string) (int, error) {

	regex := "(one|two|three|four|five|six|seven|eight|nine)|[0-9]"
	for i, line := range lines {
		lines[i] = processLine(line)
	}
	numbers, err := ProcessLines(lines, regex)
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
	createMap()
	// Read input
	filepath := "day01.txt"
	lines, err := ParseInput(filepath)
	if err != nil {
		log.Fatalf("Error while reading file: %v", err)
	}

	sol01, err := solvePart1(lines)
	if err != nil {
		log.Fatalf("Error while solving problem 1-1: %v", err)
	}
	fmt.Printf("Solution for 1-1: %d\n", sol01)

	sol02, err := solvePart2(lines)
	if err != nil {
		log.Fatalf("Error while solving problem 1-2: %v", err)
	}
	fmt.Printf("Solution for 1-2: %d\n", sol02)
}
