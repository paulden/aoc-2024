package main

import (
	"aoc-2024/internal/day01"
	"aoc-2024/internal/day02"
	"aoc-2024/internal/day03"
	"aoc-2024/internal/day04"
	"aoc-2024/internal/pkg/utils"
	"fmt"
	"strings"
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
	inputDay03 := strings.Join(utils.ReadStringsInFile("assets/day03.txt"), "")
	day03Part1 := day03.Part1(inputDay03)
	day03Part2 := day03.Part2(inputDay03)
	fmt.Println("--- Day 3: Mull It Over ---")
	fmt.Printf("Part 1: %v\n", day03Part1)
	fmt.Printf("Part 2: %v\n", day03Part2)
	inputDay04 := utils.ReadStringsInFile("assets/day04.txt")
	day04Part1 := day04.Part1(inputDay04)
	day04Part2 := day04.Part2(inputDay04)
	fmt.Println("--- Day 4: Ceres Search ---")
	fmt.Printf("Part 1: %v\n", day04Part1)
	fmt.Printf("Part 2: %v\n", day04Part2)
}
