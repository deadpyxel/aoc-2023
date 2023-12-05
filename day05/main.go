package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/deadpyxel/aoc-2023/pkg/aocutils"
)

type Almanac struct {
	seeds      []int
	seedToSoil []ConversionMap
}

type ConversionMap struct {
	sourceRangeStart int
	destRangeStart   int
	rangeSize        int
}

func parseSeeds(line string) ([]int, error) {
	_, seedNums, err := aocutils.Split(line, ":")
	if err != nil {
		return nil, err
	}
	numbers := strings.Fields(seedNums)
	nums, err := aocutils.StrToIntSlice(numbers)
	if err != nil {
		return nil, err
	}
	return nums, nil
}

func processLines(lines []string) (*Almanac, error) {
	seeds, err := parseSeeds(lines[0])
	if err != nil {
		return nil, err
	}
	// as the second line is a blank one, we can start iteration on 3rd position
	for i, line := range lines[2:] {
		switch line {
		case "seed-to-soil map:":
			fmt.Printf("Found seed-to-soil map at line %d\n", i+3)
		case "soil-to-fertilizer map:":
			fmt.Printf("Found soil-to-fertilizer map at line %d\n", i+3)
		case "fertilizer-to-water map:":
			fmt.Printf("Found fertilizer-to-water map at line %d\n", i+3)
		case "water-to-light map:":
			fmt.Printf("Found water-to-light map at line %d\n", i+3)
		case "light-to-temperature map:":
			fmt.Printf("Found light-to-temperature map at line %d\n", i+3)
		case "temperature-to-humidity map:":
			fmt.Printf("Found temperature-to-humidity map at line %d\n", i+3)
		case "humidity-to-location map:":
			fmt.Printf("Found humidity-to-location map at line %d\n", i+3)
		}
	}
	return &Almanac{seeds: seeds}, nil
}

func solvePart1(lines []string) (int, error) {
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
	almanac, _ := processLines(lines)
	fmt.Printf("Almanac data: %v\n", almanac)
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
