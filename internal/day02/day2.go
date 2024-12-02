package day02

import "aoc-2024/internal/pkg/utils"

func Part1(input [][]int) int {
	safeReports := 0
	for _, report := range input {
		isSafe, _ := IsSafe(report)
		if isSafe {
			safeReports++
		}
	}
	return safeReports
}

func Part2(input [][]int) int {
	safeReports := 0
	for _, report := range input {
		isSafe, badLevel := IsSafe(report)
		if !isSafe {
			toleratedReport1 := GenerateToleratedReport(report, badLevel-1)
			toleratedReport2 := GenerateToleratedReport(report, badLevel)
			toleratedReport3 := GenerateToleratedReport(report, badLevel+1)
			isSafe1, _ := IsSafe(toleratedReport1)
			isSafe2, _ := IsSafe(toleratedReport2)
			isSafe3, _ := IsSafe(toleratedReport3)
			isSafe = isSafe1 || isSafe2 || isSafe3
		}
		if isSafe {
			safeReports++
		}
	}
	return safeReports
}

func IsSafe(report []int) (bool, int) {
	isDecreasing := true
	isIncreasing := true
	for i := range len(report) - 1 {
		if utils.Abs(report[i]-report[i+1]) > 3 || report[i] == report[i+1] {
			return false, i
		}
		if report[i] > report[i+1] {
			isIncreasing = false
			if !isDecreasing {
				return false, i
			}

		}
		if report[i] < report[i+1] {
			isDecreasing = false
			if !isIncreasing {
				return false, i
			}
		}
	}
	return true, 0
}

func GenerateToleratedReport(report []int, indexLevelToRemove int) []int {
	var toleratedReport []int
	for i, level := range report {
		if i != indexLevelToRemove {
			toleratedReport = append(toleratedReport, level)
		}
	}
	return toleratedReport
}
