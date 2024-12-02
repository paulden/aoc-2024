package day01

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay01Part1(t *testing.T) {
	input := utils.ReadMultipleIntegersPerLineInFile("testdata/input.txt")
	expected := 11

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 1 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay01Part2(t *testing.T) {
	input := utils.ReadMultipleIntegersPerLineInFile("testdata/input.txt")
	expected := 31

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 1 / Part 2 - Expected %v, got %v", expected, result)
	}
}
