package main

import "testing"

func TestSolvePart01(t *testing.T) {
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expected := 8
	sol, err := solvePart1(lines, &GameSet{NumRed: 12, NumGreen: 13, NumBlue: 14})
	if err != nil {
		t.Errorf("Expected no errors, got %v", err)
	}
	if sol != expected {
		t.Errorf("Expected %d, got %d", expected, sol)
	}
}

func TestSolvePart02(t *testing.T) {
	lines := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	expected := 2286
	sol, err := solvePart2(lines)
	if err != nil {
		t.Errorf("Expected no errors, got %v", err)
	}
	if sol != expected {
		t.Errorf("Expected %d, got %d", expected, sol)
	}
}
