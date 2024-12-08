package day08

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay08Part1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 14

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 8 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay08Part2SimpleExample(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input2.txt")
	expected := 9

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 8 / Part 2 (with simple example) - Expected %v, got %v", expected, result)
	}
}

func TestDay08Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 34

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 8 / Part 2 - Expected %v, got %v", expected, result)
	}
}
