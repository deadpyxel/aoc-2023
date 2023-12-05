package main

import (
	"fmt"
	"log"

	"github.com/deadpyxel/aoc-2023/pkg/aocutils"
)

type Almanac struct {
}

type ConversionMap struct {
	sourceRangeStart int
	destRangeStart   int
	rangeSize        int
}

func solvePart1(lines []string) (int, error) {
	for _, line := range lines {
		fmt.Printf("%s\n", line)
	}
	return 0, nil
}

func solvePart2(lines []string) (int, error) {
	return 0, nil
}

func main() {
	lines, err := aocutils.ParseInput("day05.txt")
	if err != nil {
		log.Fatalf("Error reading input data: %v", err)
	}
	sol, err := solvePart1(lines)
	if err != nil {
		log.Fatalf("Error solving problem 5-1: %v", err)
	}
	fmt.Printf("Solution for Problem 5-1: %d\n", sol)

	sol, err = solvePart2(lines)
	if err != nil {
		log.Fatalf("Error solving problem 5-2: %v", err)
	}
	fmt.Printf("Solution for Problem 5-2: %d\n", sol)
}
