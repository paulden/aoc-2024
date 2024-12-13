package day13

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay13Part1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 480

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 13 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay13Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 875318608908

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 13 / Part 2 - Expected %v, got %v", expected, result)
	}
}
