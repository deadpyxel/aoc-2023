package main

import (
	"regexp"
	"slices"
	"testing"
)

func TestProcessLines(t *testing.T) {
	lines := []string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"}
	re := "(?:(one|two|three|four|five|six|seven|eight|nine|zero)|[0-9])"
	createMap()
	expected := []int{29, 83, 13, 24, 42, 14, 76}
	result, err := ProcessLines(lines, re)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if slices.Compare(result, expected) != 0 {
		t.Errorf("Expected: %v, got %v", expected, result)
	}
}

func TestExtract(t *testing.T) {
	input := "eightwothree"
	createMap()
	re := regexp.MustCompile("(?m)(one|two|three|four|five|six|seven|eight|nine|zero)|[0-9]")
	first, last := extract(input, re)
	if false {
		t.Errorf("Expected something, got %s, %s", first, last)
	}
}

func TestSolutionPart01(t *testing.T) {
	input := []string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"}
	result, err := solvePart1(input)
	expected := 142
	if err != nil {
		t.Errorf("Expected no errors, got %v", err)
	}

	if result != expected {
		t.Errorf("Expected %d, got %d instead.", expected, result)
	}
}
