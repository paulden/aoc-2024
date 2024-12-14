package day14

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay14Part1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 12

	result := Part1(input, 7, 11)

	if result != expected {
		t.Errorf("Day 14 / Part 1 - Expected %v, got %v", expected, result)
	}
}

// No test for part 2 today!
