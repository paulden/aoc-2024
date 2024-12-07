package day07

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	equations := ParseEquations(input)
	totalCalibrationResult := 0
	for testValue, numbers := range equations {
		possibleResults := ListPossibleResults(0, numbers)
		if slices.Contains(possibleResults, testValue) {
			totalCalibrationResult += testValue
		}
	}
	return totalCalibrationResult
}

func Part2(input []string) int {
	equations := ParseEquations(input)
	totalCalibrationResult := 0
	for testValue, numbers := range equations {
		possibleResults := ListPossibleResultsConcatenate(0, numbers, testValue)
		if slices.Contains(possibleResults, testValue) {
			totalCalibrationResult += testValue
		}
	}
	return totalCalibrationResult
}

func ListPossibleResultsConcatenate(currentValue int, numbers []int, limit int) []int {
	if len(numbers) == 0 || currentValue > limit {
		return []int{currentValue}
	}
	concatenateResult, _ := strconv.Atoi(fmt.Sprintf("%d%d", currentValue, numbers[0]))
	return append(
		ListPossibleResultsConcatenate(concatenateResult, numbers[1:], limit),
		append(
			ListPossibleResultsConcatenate(currentValue+numbers[0], numbers[1:], limit),
			ListPossibleResultsConcatenate(currentValue*numbers[0], numbers[1:], limit)...,
		)...,
	)
}

func ListPossibleResults(currentValue int, numbers []int) []int {
	if len(numbers) == 0 {
		return []int{currentValue}
	}
	return append(
		ListPossibleResults(currentValue+numbers[0], numbers[1:]),
		ListPossibleResults(currentValue*numbers[0], numbers[1:])...,
	)
}

func ParseEquations(input []string) map[int][]int {
	equations := make(map[int][]int)
	for _, line := range input {
		equation := strings.Split(line, ": ")
		testValue, _ := strconv.Atoi(equation[0])
		var remainingNumbers []int
		for _, remainingNumber := range strings.Split(equation[1], " ") {
			parsedNumber, _ := strconv.Atoi(remainingNumber)
			remainingNumbers = append(remainingNumbers, parsedNumber)
		}
		equations[testValue] = remainingNumbers
	}
	return equations
}
