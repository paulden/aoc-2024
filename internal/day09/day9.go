package day09

import (
	"slices"
	"strconv"
)

type sector struct {
	start, end int
}

func Part1(input []string) int {
	diskMap := GetDiskMap(input[0])
	nextEmptySector := 0
	for diskMap[nextEmptySector] >= 0 {
		nextEmptySector++
	}
	for i := len(diskMap) - 1; i > nextEmptySector; i-- {
		if diskMap[i] >= 0 {
			diskMap[nextEmptySector] = diskMap[i]
			diskMap[i] = -1
		}
		for diskMap[nextEmptySector] >= 0 {
			nextEmptySector++
		}
	}
	return ComputeChecksum(diskMap)
}

func Part2(input []string) int {
	diskMap := GetDiskMap(input[0])
	emptySectors := ListEmptySectors(diskMap)
	i := len(diskMap) - 1
	for i >= 0 {
		if diskMap[i] >= 0 {
			fileID := diskMap[i]
			currentSectorEnd := i
			for i >= 0 && diskMap[i] == fileID {
				i--
			}
			currentSectorStart := i + 1
			sectorSize := currentSectorEnd - currentSectorStart
			for _, emptySector := range emptySectors {
				emptySectorSize := emptySector.end - emptySector.start
				if emptySector.end < currentSectorStart && emptySectorSize >= sectorSize {
					for c := emptySector.start; c <= emptySector.start+sectorSize; c++ {
						diskMap[c] = fileID
					}
					for c := currentSectorStart; c <= currentSectorEnd; c++ {
						diskMap[c] = -1
					}
					emptySectors = ListEmptySectors(diskMap)
					break
				}
			}
		} else {
			i--
		}
	}
	return ComputeChecksum(diskMap)
}

func GetDiskMap(input string) []int {
	var diskMap []int
	blockID := 0
	for i, fileBlock := range input {
		blockSize, _ := strconv.Atoi(string(fileBlock))
		if i%2 == 0 {
			diskMap = append(diskMap, slices.Repeat([]int{blockID}, blockSize)...)
			blockID++
		} else {
			diskMap = append(diskMap, slices.Repeat([]int{-1}, blockSize)...)
		}
	}
	return diskMap
}

func ListEmptySectors(diskMap []int) []sector {
	var emptySectors []sector
	i := 0
	for i < len(diskMap) {
		if diskMap[i] == -1 {
			var emptySector sector
			emptySector.start = i
			for i < len(diskMap) && diskMap[i] < 0 {
				i++
			}
			emptySector.end = i - 1
			emptySectors = append(emptySectors, emptySector)
		} else {
			i++
		}
	}
	return emptySectors
}

func ComputeChecksum(diskMap []int) int {
	checksum := 0
	for i, block := range diskMap {
		if block < 0 {
			continue
		}
		checksum += i * block
	}
	return checksum
}
