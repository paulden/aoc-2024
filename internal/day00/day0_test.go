package day00

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay00Part1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 1

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 0 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay00Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 1

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 0 / Part 2 - Expected %v, got %v", expected, result)
	}
}
