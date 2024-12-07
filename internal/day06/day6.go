package day06

import (
	"fmt"
)

type coordinates struct {
	x, y int
}

type guardPosition struct {
	x, y      int
	direction string
}

func (guard *guardPosition) IsLeaving(labMap [][]int) bool {
	if guard.direction == "N" && guard.x == 0 {
		return true
	}
	if guard.direction == "S" && guard.x == len(labMap)-1 {
		return true
	}
	if guard.direction == "W" && guard.y == 0 {
		return true
	}
	if guard.direction == "E" && guard.y == len(labMap[0])-1 {
		return true
	}
	return false
}

func (guard *guardPosition) Turn() {
	switch guard.direction {
	case "N":
		guard.direction = "E"
	case "E":
		guard.direction = "S"
	case "S":
		guard.direction = "W"
	case "W":
		guard.direction = "N"
	}
}

func (guard *guardPosition) MoveForward() {
	switch guard.direction {
	case "N":
		guard.x--
	case "E":
		guard.y++
	case "S":
		guard.x++
	case "W":
		guard.y--
	}
}

func (guard *guardPosition) IsStuck(labMap [][]int) bool {
	switch guard.direction {
	case "N":
		return labMap[guard.x-1][guard.y] == 1
	case "E":
		return labMap[guard.x][guard.y+1] == 1
	case "S":
		return labMap[guard.x+1][guard.y] == 1
	case "W":
		return labMap[guard.x][guard.y-1] == 1
	}
	return false
}

func Part1(input []string) int {
	labMap, initialGuard := GetMapAndGuardPosition(input)
	visitedTiles := ListVisitedTiles(labMap, initialGuard)
	return len(visitedTiles)
}

func Part2(input []string) int {
	labMap, initialGuard := GetMapAndGuardPosition(input)
	potentialObstacles := ListVisitedTiles(labMap, initialGuard)
	delete(potentialObstacles, coordinates{initialGuard.x, initialGuard.y})
	obstacles := 0

	for potentialObstacle := range potentialObstacles {
		mapWithObstacle := AddObstacleToLabMad(labMap, potentialObstacle)
		previousGuardPositions := make(map[guardPosition]bool)
		guard := initialGuard
		previousGuardPositions[guard] = true
		for !guard.IsLeaving(mapWithObstacle) {
			for guard.IsStuck(mapWithObstacle) {
				guard.Turn()
			}
			guard.MoveForward()

			if previousGuardPositions[guard] {
				obstacles++
				break
			}

			previousGuardPositions[guard] = true
		}
	}

	return obstacles
}

func GetMapAndGuardPosition(input []string) ([][]int, guardPosition) {
	labMap := make([][]int, len(input))
	var guard guardPosition
	for i, line := range input {
		for j, tile := range line {
			if tile == '^' {
				guard = guardPosition{i, j, "N"}
				labMap[i] = append(labMap[i], 0)
			}
			if tile == '#' {
				labMap[i] = append(labMap[i], 1)
			}
			if tile == '.' {
				labMap[i] = append(labMap[i], 0)
			}
		}
	}
	return labMap, guard
}

func ListVisitedTiles(labMap [][]int, initialGuardPosition guardPosition) map[coordinates]bool {
	visitedTiles := make(map[coordinates]bool)
	guard := initialGuardPosition
	visitedTiles[coordinates{guard.x, guard.y}] = true
	for !guard.IsLeaving(labMap) {
		for guard.IsStuck(labMap) {
			guard.Turn()
		}
		guard.MoveForward()
		visitedTiles[coordinates{guard.x, guard.y}] = true
	}
	return visitedTiles
}

func PrettyPrintLabMap(labMap [][]int, visitedTiles map[coordinates]bool) {
	for i := range labMap {
		for j := range labMap[0] {
			if visitedTiles[coordinates{i, j}] {
				fmt.Printf("X")
			} else if labMap[i][j] == 1 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
}

func AddObstacleToLabMad(labMap [][]int, obstacle coordinates) [][]int {
	newLabMap := make([][]int, len(labMap))
	for i := range labMap {
		newLabMap[i] = make([]int, len(labMap[i]))
		copy(newLabMap[i], labMap[i])
	}
	newLabMap[obstacle.x][obstacle.y] = 1
	return newLabMap
}
