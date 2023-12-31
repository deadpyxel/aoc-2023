package main

import "testing"

func TestSolverPart01(t *testing.T) {
	lines := []string{
		"seeds: 79 14 55 13",
		"\n",
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
		"\n",
		"soil-to-fertilizer map:",
		"0 15 37",
		"37 52 2",
		"39 0 15",
		"\n",
		"fertilizer-to-water map:",
		"49 53 8",
		"0 11 42",
		"42 0 7",
		"57 7 4",
		"\n",
		"water-to-light map:",
		"88 18 7",
		"18 25 70",
		"\n",
		"light-to-temperature map:",
		"45 77 23",
		"81 45 19",
		"68 64 13",
		"\n",
		"temperature-to-humidity map:",
		"0 69 1",
		"1 0 69",
		"\n",
		"humidity-to-location map:",
		"60 56 37",
		"56 93 4",
	}
	sol, _ := solvePart1(lines)
	expected := -1
	if sol != expected {
		t.Errorf("Expected %d, got %d\n", expected, sol)
	}
}
