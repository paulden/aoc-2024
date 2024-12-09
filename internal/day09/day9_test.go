package day09

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay09Part1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 1928

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 9 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay09Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 2858

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 9 / Part 2 - Expected %v, got %v", expected, result)
	}
}
