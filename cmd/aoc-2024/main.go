package main

import (
	"aoc-2024/internal/day01"
	"aoc-2024/internal/day02"
	"aoc-2024/internal/pkg/utils"
	"fmt"
)

func main() {
	inputDay01 := utils.ReadMultipleIntegersPerLineInFile("assets/day01.txt")
	day01Part1 := day01.Part1(inputDay01)
	day01Part2 := day01.Part2(inputDay01)
	fmt.Println("--- Day 1: Historian Hysteria ---")
	fmt.Printf("Part 1: %v\n", day01Part1)
	fmt.Printf("Part 2: %v\n", day01Part2)
	inputDay02 := utils.ReadMultipleIntegersPerLineInFile("assets/day02.txt")
	day02Part1 := day02.Part1(inputDay02)
	day02Part2 := day02.Part2(inputDay02)
	fmt.Println("--- Day 2: Red-Nosed Reports ---")
	fmt.Printf("Part 1: %v\n", day02Part1)
	fmt.Printf("Part 2: %v\n", day02Part2)
}
