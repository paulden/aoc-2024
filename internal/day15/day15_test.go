package day15

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay15Part1Example1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input1.txt")
	expected := 2028

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 15 / Part 1 / Example 1 - Expected %v, got %v", expected, result)
	}
}


func TestDay15Part1Example2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input1.txt")
	expected := 2028

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 15 / Part 1 / Example 2 - Expected %v, got %v", expected, result)
	}
}

func TestDay15Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input.txt")
	expected := 1

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 0 / Part 2 - Expected %v, got %v", expected, result)
	}
}
