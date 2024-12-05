package day05

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay05Part1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 143

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 5 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay05Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 123

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 5 / Part 2 - Expected %v, got %v", expected, result)
	}
}
