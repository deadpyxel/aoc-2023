package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/deadpyxel/aod-2023/pkg/aocutils"
)

type GameSet struct {
	NumRed   int
	NumGreen int
	NumBlue  int
}

type Game struct {
	ID       int
	GameSets []GameSet
}

func processLine(line string) (*Game, error) {
	gameData := strings.Split(line, ": ")
	gameId, err := strconv.Atoi(gameData[0][5:])
	if err != nil {
		return nil, err
	}
	sets := []GameSet{}
	gameSets := strings.Split(gameData[1], "; ")
	for _, gameSet := range gameSets {
		setColors := strings.Split(gameSet, ", ")
		var amountR, amountG, amountB int
		for _, colorInfo := range setColors {
			colorData := strings.Split(colorInfo, " ")
			amount, err := strconv.Atoi(colorData[0])
			if err != nil {
				return nil, err
			}
			switch colorData[1] {
			case "red":
				amountR = amount
			case "green":
				amountG = amount
			case "blue":
				amountB = amount
			}
		}
		set := GameSet{NumRed: amountR, NumGreen: amountG, NumBlue: amountB}
		sets = append(sets, set)
	}
	return &Game{ID: gameId, GameSets: sets}, nil
}

func solvePart1(lines []string, constraint *GameSet) (int, error) {
	var gameIds []int
	for _, line := range lines {
		isInvalid := false
		// Extract Game info from line
		g, err := processLine(line)
		if err != nil {
			return 0, err
		}
		// Check each of the sets in that game, if any are above the constraints, discard that game
		for _, set := range g.GameSets {
			if set.NumRed > constraint.NumRed || set.NumGreen > constraint.NumGreen || set.NumBlue > constraint.NumBlue {
				isInvalid = true
				break
			}
		}
		// If the game is considered valid, add it to the pool
		if !isInvalid {
			gameIds = append(gameIds, g.ID)
		}
	}
	return aocutils.Sum(gameIds), nil
}

func solvePart2(lines []string) (int, error) {
	totalProd := 0
	for _, line := range lines {
		// Extract Game info from line
		g, err := processLine(line)
		if err != nil {
			return 0, err
		}
		minR, minG, minB := 0, 0, 0
		// For each subset of a game, checks if our current minimum is below the needed cubes for this set
		// if that is the case, update the minimums
		for _, set := range g.GameSets {
			if set.NumRed > minR {
				minR = set.NumRed
			}
			if set.NumGreen > minG {
				minG = set.NumGreen
			}
			if set.NumBlue > minB {
				minB = set.NumBlue
			}
		}
		totalProd += aocutils.Product(minR, minG, minB)
	}
	return totalProd, nil
}
func main() {
	lines, err := aocutils.ParseInput("day02.txt")
	if err != nil {
		log.Fatalf("Error reading input data: %v", err)
	}
	sol, err := solvePart1(lines, &GameSet{NumRed: 12, NumGreen: 13, NumBlue: 14})
	if err != nil {
		log.Fatalf("Error solving problem 2-1: %v", err)
	}
	fmt.Printf("Solution for Problem 2-1: %d\n", sol)

	sol, err = solvePart2(lines)
	if err != nil {
		log.Fatalf("Error solving problem 2-2: %v", err)
	}
	fmt.Printf("Solution for Problem 2-2: %d\n", sol)
}
