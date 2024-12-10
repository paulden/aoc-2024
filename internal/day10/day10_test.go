package day10

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
)

func TestDay10Part1Input2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input2.txt")
	expected := 2

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 10 / Part 1 / Input 2 - Expected %v, got %v", expected, result)
	}
}

func TestDay10Part1Input3(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input3.txt")
	expected := 4

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 10 / Part 1 / Input 3 - Expected %v, got %v", expected, result)
	}
}

func TestDay10Part1Input4(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input4.txt")
	expected := 3

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 10 / Part 1 / Input 4 - Expected %v, got %v", expected, result)
	}
}

func TestDay10Part1Input5(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input5.txt")
	expected := 36

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 10 / Part 1 / Input 5 - Expected %v, got %v", expected, result)
	}
}

func TestDay10Part2(t *testing.T) {
	input := utils.ReadStringsInFile("testdata/input5.txt")
	expected := 81

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 10 / Part 2 - Expected %v, got %v", expected, result)
	}
}
