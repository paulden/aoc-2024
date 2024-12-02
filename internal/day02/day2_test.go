package day02

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay02Part1(t *testing.T) {
	input := utils.ReadMultipleIntegersPerLineInFile("testdata/input.txt")
	expected := 2

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 2 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay02Part2(t *testing.T) {
	input := utils.ReadMultipleIntegersPerLineInFile("testdata/input.txt")
	expected := 4

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 2 / Part 2 - Expected %v, got %v", expected, result)
	}
}
