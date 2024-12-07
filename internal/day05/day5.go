package day05

import (
	"slices"
	"strconv"
	"strings"
)

func Part1(input []string) int {
	splitIndex := slices.Index(input, "")
	pageOrderingRules := GetPageOrderingRules(input[:splitIndex])
	pagesUpdates := GetPageNumbers(input[splitIndex+1:])

	middlePageSum := 0

	for _, pagesUpdate := range pagesUpdates {
		if IsValid(pagesUpdate, pageOrderingRules) {
			middlePageSum += pagesUpdate[len(pagesUpdate)/2]
		}
	}
	return middlePageSum
}

func Part2(input []string) int {
	splitIndex := slices.Index(input, "")
	pageOrderingRules := GetPageOrderingRules(input[:splitIndex])
	pagesUpdates := GetPageNumbers(input[splitIndex+1:])

	middlePageSum := 0

	for _, pagesUpdate := range pagesUpdates {
		if !IsValid(pagesUpdate, pageOrderingRules) {
			slices.SortFunc(pagesUpdate, func(a, b int) int {
				if slices.Contains(pageOrderingRules[a], b) {
					return -1
				}
				return 1
			})
			middlePageSum += pagesUpdate[len(pagesUpdate)/2]
		}
	}
	return middlePageSum
}

func GetPageOrderingRules(rules []string) map[int][]int {
	pageOrderingRules := make(map[int][]int)
	for _, rule := range rules {
		pair := strings.Split(rule, "|")
		before, _ := strconv.Atoi(pair[0])
		after, _ := strconv.Atoi(pair[1])
		pageOrderingRules[before] = append(pageOrderingRules[before], after)
	}
	return pageOrderingRules
}

func GetPageNumbers(updates []string) [][]int {
	pagesUpdates := make([][]int, len(updates))
	for i, rule := range updates {
		pages := strings.Split(rule, ",")
		for _, page := range pages {
			parsedPage, _ := strconv.Atoi(page)
			pagesUpdates[i] = append(pagesUpdates[i], parsedPage)
		}
	}
	return pagesUpdates
}

func IsValid(pagesUpdate []int, pagesOrderingRules map[int][]int) bool {
	for i, page := range pagesUpdate {
		for _, pageBefore := range pagesUpdate[0:i] {
			if slices.Contains(pagesOrderingRules[page], pageBefore) {
				return false
			}
		}
	}
	return true
}
