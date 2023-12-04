package main

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/deadpyxel/aoc-2023/pkg/aocutils"
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

// addCopies adds copies of cards to the `cardCopies` map based on the number of `matches` and the current position `currPos`.
func addCopies(matches int, currPos int, cardCopies map[int]int) {

	offsetPos := currPos + matches

	for k := currPos + 1; k <= offsetPos; k++ {
		cardCopies[k]++
	}
}

func solvePart2(lines []string) (int, error) {
	// cardCopies will keep track of how many copies we have of each scratch card
	cardCopies := make(map[int]int, len(lines))
	for i := range lines {
		cardCopies[i] = 1
	}
	// matchIdx will work as a lookup table to avoid recalculations
	matchIdx := make(map[int]int, len(lines))
	for i := range lines {
		matchIdx[i] = -1
	}
	for i, line := range lines {
		c := processLine(line)
		cardMatches := 0

		// For each card (original or copy) check for matches
		for j := 0; j < cardCopies[i]; j++ {
			// if this card already has been processed, use the lookup table values
			if matchIdx[i] != -1 {
				addCopies(matchIdx[i], i, cardCopies)
				continue
			}
			for _, num := range c.CardNumbers {
				if slices.Contains(c.WinningNumbers, num) {
					cardMatches++
				}
			}
			// add this number of matches to the cache
			matchIdx[i] = cardMatches

			// add the copies
			addCopies(cardMatches, i, cardCopies)

			// reset match count
			cardMatches = 0
		}
	}

	totalPoints := 0
	for _, copyCount := range cardCopies {
		totalPoints += copyCount
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

	sol, err = solvePart2(lines)
	if err != nil {
		log.Fatalf("Error solving problem 4-2: %v", err)
	}
	fmt.Printf("Solution for Problem 4-2: %d\n", sol)
}
