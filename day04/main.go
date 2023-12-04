package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/deadpyxel/aod-2023/pkg/aocutils"
)

type Card struct {
	WinningNumbers []string
	CardNumbers    []string
}

func processLine(line string) *Card {
	parts := strings.Split(line, "|")
	winningNum := strings.Fields(parts[0])
	cardNum := strings.Fields(parts[1])
	return &Card{WinningNumbers: winningNum[2:], CardNumbers: cardNum}
}

func solvePart1(lines []string) (int, error) {
	totalPoints := 0
	for _, line := range lines {
		cardPoints := 0
		c := processLine(line)
		for _, num := range c.CardNumbers {
			if slices.Contains(c.WinningNumbers, num) {
				if cardPoints == 0 {
					cardPoints = 1
					continue
				}
				cardPoints *= 2
			}
		}
		totalPoints += cardPoints
	}
	return totalPoints, nil
}

func main() {
	lines, err := aocutils.ParseInput("day04.txt")
	if err != nil {
		log.Fatalf("Error reading input data: %v", err)
	}
	sol, err := solvePart1(lines)
	if err != nil {
		log.Fatalf("Error solving problem 4-1: %v", err)
	}
	fmt.Printf("Solution for Problem 4-1: %d\n", sol)

	// sol, err = solvePart2(lines)
	// if err != nil {
	// 	log.Fatalf("Error solving problem 4-2: %v", err)
	// }
	// fmt.Printf("Solution for Problem 4-2: %d\n", sol)
}
