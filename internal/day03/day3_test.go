package day03

import (
	"aoc-2024/internal/pkg/utils"
	"testing"
	"strings"
)

func TestDay03Part1(t *testing.T) {
	input := strings.Join(utils.ReadStringsInFile("testdata/input.txt"), "")
	expected := 161

	result := Part1(input)

	if result != expected {
		t.Errorf("Day 3 / Part 1 - Expected %v, got %v", expected, result)
	}
}

func TestDay03Part2(t *testing.T) {
	input := strings.Join(utils.ReadStringsInFile("testdata/input2.txt"), "")
	expected := 48

	result := Part2(input)

	if result != expected {
		t.Errorf("Day 3 / Part 2 - Expected %v, got %v", expected, result)
	}
}
