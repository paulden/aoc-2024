package day11

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay11Part1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 55312

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 11 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay11Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 65601038650482

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 11 / Part 2 - Expected %v, got %v", expected, result)
	}
}
