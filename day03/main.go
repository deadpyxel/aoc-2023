package main

import (
	"fmt"
	"log"
	"unicode"

	"github.com/deadpyxel/aod-2023/pkg/aocutils"
)

func convertToSchematic(lines []string) [][]rune {
	schematic := make([][]rune, len(lines))
	for i, line := range lines {
		schematic[i] = make([]rune, len(line))
		for j, ch := range line {
			schematic[i][j] = ch
		}
	}
	return schematic
}

func solvePart1(schematic [][]rune) (int, error) {
	var partNumbers []int
	for i, line := range schematic {
		for j, ch := range line {
			// If the character is not a dot or a number, it is a symbol
			if !(unicode.IsNumber(ch) || string(ch) == ".") {
				fmt.Printf("Found symbol [%s] at (%d, %d)\n", string(ch), i, j)
				if i > 0 && i < len(schematic) {

				}
			}
		}
	}

	return aocutils.Sum(partNumbers), nil
}

func main() {

	lines, err := aocutils.ParseInput("day03.txt")
	if err != nil {
		log.Fatalf("Error reading input data: %v\n", err)
	}
	schematic := convertToSchematic(lines)
	fmt.Printf(" Schematic:\n%v\n", schematic)
	sol, err := solvePart1(schematic)
	if err != nil {
		log.Fatalf("Found error solving Problem 3-1: %v\n", err)
	}
	fmt.Printf(" Solution for problem 3-1: %d\n", sol)
}
