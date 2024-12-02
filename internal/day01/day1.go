package day01

import (
	"aoc-2024/internal/pkg/utils"
	"slices"
)

func Part1(input [][]int) int {
	leftList, rightList := GetLists(input)

	slices.Sort(leftList)
	slices.Sort(rightList)

	totalDistance := 0

	for i := range input {
		totalDistance += utils.Abs(leftList[i] - rightList[i])
	}

	return totalDistance
}

func Part2(input [][]int) int {
	leftList, rightList := GetLists(input)

	rightListOccurrences := make(map[int]int)

	for _, locationID := range rightList {
		rightListOccurrences[locationID]++
	}

	similarityScore := 0

	for _, locationID := range leftList {
		if occurrences, ok := rightListOccurrences[locationID]; ok {
			similarityScore += locationID * occurrences
		}
	}

	return similarityScore
}

func GetLists(input [][]int) ([]int, []int) {
	var leftList []int
	var rightList []int

	for _, locationIDs := range input {
		leftList = append(leftList, locationIDs[0])
		rightList = append(rightList, locationIDs[1])
	}

	return leftList, rightList
}
