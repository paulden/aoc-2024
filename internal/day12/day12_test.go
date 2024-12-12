package day12

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay12Part1Example1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input1.txt")
	expected := 140

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 12 / Part 1 / Example 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay12Part1Example2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input2.txt")
	expected := 772

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 12 / Part 1 / Example 2 - Expected %v, got %v", expected, result)
	}
}

func TestDay12Part1Example3(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input3.txt")
	expected := 1930

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 12 / Part 1 / Example 3 - Expected %v, got %v", expected, result)
	}
}

func TestDay12Part2Example1(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input1.txt")
	expected := 80

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 12 / Part 2 / Example 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay12Part2Example3(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input3.txt")
	expected := 1206

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 12 / Part 2 / Example 3 - Expected %v, got %v", expected, result)
	}
}
