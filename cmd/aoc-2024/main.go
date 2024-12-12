package main

import (
	"aoc-2024/internal/day01"
	"aoc-2024/internal/day02"
	"aoc-2024/internal/day03"
	"aoc-2024/internal/day04"
	"aoc-2024/internal/day05"
	"aoc-2024/internal/day06"
	"aoc-2024/internal/day07"
	"aoc-2024/internal/day08"
	"aoc-2024/internal/day09"
	"aoc-2024/internal/day10"
	"aoc-2024/internal/day11"
	"aoc-2024/internal/day12"
	"aoc-2024/internal/pkg/utils"
	"flag"
	"fmt"
	"strings"
)

func main() {
	day := flag.Int("day", -1, "Day to run")
	flag.Parse()

	switch *day {
	case 1:
		inputDay01 := utils.ReadMultipleIntegersPerLineInFile("assets/day01.txt")
		day01Part1 := day01.Part1(inputDay01)
		day01Part2 := day01.Part2(inputDay01)
		fmt.Println("--- Day 1: Historian Hysteria ---")
		fmt.Printf("Part 1: %v\n", day01Part1)
		fmt.Printf("Part 2: %v\n", day01Part2)
	case 2:
		inputDay02 := utils.ReadMultipleIntegersPerLineInFile("assets/day02.txt")
		day02Part1 := day02.Part1(inputDay02)
		day02Part2 := day02.Part2(inputDay02)
		fmt.Println("--- Day 2: Red-Nosed Reports ---")
		fmt.Printf("Part 1: %v\n", day02Part1)
		fmt.Printf("Part 2: %v\n", day02Part2)
	case 3:
		inputDay03 := strings.Join(utils.ReadStringsInFile("assets/day03.txt"), "")
		day03Part1 := day03.Part1(inputDay03)
		day03Part2 := day03.Part2(inputDay03)
		fmt.Println("--- Day 3: Mull It Over ---")
		fmt.Printf("Part 1: %v\n", day03Part1)
		fmt.Printf("Part 2: %v\n", day03Part2)
	case 4:
		inputDay04 := utils.ReadStringsInFile("assets/day04.txt")
		day04Part1 := day04.Part1(inputDay04)
		day04Part2 := day04.Part2(inputDay04)
		fmt.Println("--- Day 4: Ceres Search ---")
		fmt.Printf("Part 1: %v\n", day04Part1)
		fmt.Printf("Part 2: %v\n", day04Part2)
	case 5:
		inputDay05 := utils.ReadStringsInFile("assets/day05.txt")
		day05Part1 := day05.Part1(inputDay05)
		day05Part2 := day05.Part2(inputDay05)
		fmt.Println("--- Day 5: Print Queue ---")
		fmt.Printf("Part 1: %v\n", day05Part1)
		fmt.Printf("Part 2: %v\n", day05Part2)
	case 6:
		inputDay06 := utils.ReadStringsInFile("assets/day06.txt")
		day06Part1 := day06.Part1(inputDay06)
		day06Part2 := day06.Part2(inputDay06)
		fmt.Println("--- Day 6: Guard Gallivant ---")
		fmt.Printf("Part 1: %v\n", day06Part1)
		fmt.Printf("Part 2: %v\n", day06Part2)
	case 7:
		inputDay07 := utils.ReadStringsInFile("assets/day07.txt")
		day07Part1 := day07.Part1(inputDay07)
		day07Part2 := day07.Part2(inputDay07)
		fmt.Println("--- Day 7: Bridge Repair ---")
		fmt.Printf("Part 1: %v\n", day07Part1)
		fmt.Printf("Part 2: %v\n", day07Part2)
	case 8:
		inputDay08 := utils.ReadStringsInFile("assets/day08.txt")
		day08Part1 := day08.Part1(inputDay08)
		day08Part2 := day08.Part2(inputDay08)
		fmt.Println("--- Day 8: Resonant Collinearity ---")
		fmt.Printf("Part 1: %v\n", day08Part1)
		fmt.Printf("Part 2: %v\n", day08Part2)
	case 9:
		inputDay09 := utils.ReadStringsInFile("assets/day09.txt")
		day09Part1 := day09.Part1(inputDay09)
		day09Part2 := day09.Part2(inputDay09)
		fmt.Println("--- Day 9: Disk Fragmenter ---")
		fmt.Printf("Part 1: %v\n", day09Part1)
		fmt.Printf("Part 2: %v\n", day09Part2)
	case 10:
		inputDay10 := utils.ReadStringsInFile("assets/day10.txt")
		day10Part1 := day10.Part1(inputDay10)
		day10Part2 := day10.Part2(inputDay10)
		fmt.Println("--- Day 10: Hoof It ---")
		fmt.Printf("Part 1: %v\n", day10Part1)
		fmt.Printf("Part 2: %v\n", day10Part2)
	case 11:
		inputDay11 := utils.ReadStringsInFile("assets/day11.txt")
		day11Part1 := day11.Part1(inputDay11)
		day11Part2 := day11.Part2(inputDay11)
		fmt.Println("--- Day 11: Plutonian Pebbles ---")
		fmt.Printf("Part 1: %v\n", day11Part1)
		fmt.Printf("Part 2: %v\n", day11Part2)
	case 12:
		inputDay12 := utils.ReadStringsInFile("assets/day12.txt")
		day12Part1 := day12.Part1(inputDay12)
		day12Part2 := day12.Part2(inputDay12)
		fmt.Println("--- Day 12: Garden Groups ---")
		fmt.Printf("Part 1: %v\n", day12Part1)
		fmt.Printf("Part 2: %v\n", day12Part2)

	}
}
