package day06

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay06Part1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 41

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 6 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay06Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 6

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 6 / Part 2 - Expected %v, got %v", expected, result)
	}
}

func TestDay06Part2Input2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input2.txt")
	expected := 19

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 6 / Part 2 with second input - Expected %v, got %v", expected, result)
	}
}
