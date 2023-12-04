package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"unicode"

	"github.com/deadpyxel/aoc-2023/pkg/aocutils"
)

func convertToSchematic(lines []string) [][]rune {
	schematic := make([][]rune, len(lines))
	for i, line := range lines {
		schematic[i] = []rune(line)
	}
	return schematic
}

func isSymbol(r rune) bool {
	return !(unicode.IsNumber(r) || string(r) == ".")
}

func solvePart2(schematic [][]rune) (int, error) {
	total := 0
	for i, line := range schematic {
		j := 0
		for j < len(line) {
			ch := line[j]
			// If the character is not '*', skip this character
			if ch != '*' {
				j++
				continue
			}

			// Check the characters around the current symbol for part numbers
			partNumbers := []int{}
			for di := -1; di <= 1; di++ {
				for dj := -1; dj <= 1; dj++ {
					if di == 0 && dj == 0 || i+di < 0 || i+di >= len(schematic) || j+dj < 0 || j+dj >= len(line) {
						continue
					}
					if unicode.IsDigit(schematic[i+di][j+dj]) {
						// Get the full number
						k := j + dj
						for k < len(line) && unicode.IsDigit(line[k]) {
							k++
						}
						num, _ := strconv.Atoi(string(line[j+dj : k]))
						partNumbers = append(partNumbers, num)
					}
				}
			}

			// If exactly two part numbers are adjacent to the symbol, it's a gear
			if len(partNumbers) == 2 {
				total += partNumbers[0] * partNumbers[1]
			}
			j++
		}
	}

	return total, nil
}

func solvePart1(schematic [][]rune) (int, error) {
	var partNumbers []int
	numRegex := regexp.MustCompile("[0-9]+")
	for i, line := range schematic {
		numPos := numRegex.FindAllStringIndex(string(line), -1)
		for _, pos := range numPos {
			numStr := string(line[pos[0]:pos[1]])
			num, _ := strconv.Atoi(numStr)
			if isAdjacentToSymbol(schematic, i, pos[0], len(numStr)) {
				partNumbers = append(partNumbers, num)
			}
		}
	}

	return aocutils.Sum(partNumbers), nil
}

func isAdjacentToSymbol(schematic [][]rune, i, j, length int) bool {
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= length; dj++ {
			if isOutOfBounds(schematic, i+di, j+dj) {
				continue
			}
			if isSymbol(schematic[i+di][j+dj]) {
				return true
			}
		}
	}
	return false
}

func isOutOfBounds(schematic [][]rune, i, j int) bool {
	return i < 0 || i >= len(schematic) || j < 0 || j >= len(schematic[i])
}

func main() {

	lines, err := aocutils.ParseInput("day03.txt")
	if err != nil {
		log.Fatalf("Error reading input data: %v\n", err)
	}
	schematic := convertToSchematic(lines)
	sol, err := solvePart1(schematic)
	if err != nil {
		log.Fatalf("Found error solving Problem 3-1: %v\n", err)
	}
	fmt.Printf(" Solution for problem 3-1: %d\n", sol)

	sol, err = solvePart2(schematic)
	if err != nil {
		log.Fatalf("Found error solving Problem 3-2: %v\n", err)
	}
	fmt.Printf(" Solution for problem 3-2: %d\n", sol)
}
