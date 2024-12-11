package day11

import (
	"strconv"
	"strings"
)

func Part1(input []string) int {
	stones := GetStonesArrangement(input[0])
	for i := 0; i < 25; i++ {
		stones = GetNextStonesArrangement(stones)
	}
	return len(stones)
}

func Part2(input []string) int {
	stones := CountStones(input[0])
	for i := 0; i < 75; i++ {
		stones = GetNextStonesCount(stones)
	}
	totalStonesCount := 0
	for _, count := range stones {
		totalStonesCount += count
	}
	return totalStonesCount
}

func GetStonesArrangement(input string) []int {
	stones := strings.Split(input, " ")
	stonesArrangement := make([]int, len(stones))
	for i, stone := range stones {
		stoneInteger, _ := strconv.Atoi(stone)
		stonesArrangement[i] = stoneInteger
	}
	return stonesArrangement
}

func CountStones(input string) map[int]int {
	stones := strings.Split(input, " ")
	stonesCount := make(map[int]int)
	for _, stone := range stones {
		stoneInteger, _ := strconv.Atoi(stone)
		stonesCount[stoneInteger]++
	}
	return stonesCount
}

func GetNextStonesArrangement(stones []int) []int {
	var nextStonesArrangement []int
	for _, stone := range stones {
		stringStone := strconv.Itoa(stone)
		if stone == 0 {
			nextStonesArrangement = append(nextStonesArrangement, 1)
		} else if len(stringStone)%2 == 0 {
			half := len(stringStone) / 2
			firstHalf, _ := strconv.Atoi(stringStone[:half])
			secondHalf, _ := strconv.Atoi(stringStone[half:])
			nextStonesArrangement = append(nextStonesArrangement, []int{firstHalf, secondHalf}...)
		} else {
			nextStonesArrangement = append(nextStonesArrangement, stone*2024)
		}
	}
	return nextStonesArrangement
}

func GetNextStonesCount(stones map[int]int) map[int]int {
	nextStonesCount := make(map[int]int)
	for stone, count := range stones {
		stringStone := strconv.Itoa(stone)
		if stone == 0 {
			nextStonesCount[1] += count
		} else if len(stringStone)%2 == 0 {
			half := len(stringStone) / 2
			firstHalf, _ := strconv.Atoi(stringStone[:half])
			secondHalf, _ := strconv.Atoi(stringStone[half:])
			nextStonesCount[firstHalf] += count
			nextStonesCount[secondHalf] += count
		} else {
			nextStonesCount[stone*2024] += count
		}
	}
	return nextStonesCount
}
