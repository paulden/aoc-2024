package day07

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay07Part1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 3749

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 7 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay07Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 11387

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 7 / Part 2 - Expected %v, got %v", expected, result)
	}
}
