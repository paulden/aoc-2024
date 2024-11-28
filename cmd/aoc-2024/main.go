package main

import (
	"aoc-2024/internal/day01"
	"aoc-2024/internal/pkg/utils"
	"fmt"
)

func main() {
	input := utils.ReadIntegersInFile("assets/day1.txt")
	day01Part1 := day01.Part1(input)
	day01Part2 := day01.Part2(input)
	fmt.Println("--- Day 1: Title ---")
	fmt.Printf("Part 1: %v\n", day01Part1)
	fmt.Printf("Part 2: %v\n", day01Part2)
}
