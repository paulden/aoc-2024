package day04

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay04Part1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 18

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 4 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay04Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 9

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 4 / Part 2 - Expected %v, got %v", expected, result)
	}
}
