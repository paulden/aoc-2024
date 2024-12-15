package day15

import (
	"fmt"
	"slices"
)

type coordinates struct {
	x, y int
}

type robot struct {
	position coordinates
}

func Part1(input []string) int {
	splitIndex := slices.Index(input, "")
	warehouse, robot := GetInitialWarehouse(input[:splitIndex])
	movements := GetMovements(input[splitIndex+1:])

	for _, movement := range movements {
		blocksToMove := 1
		switch movement {
		case ">":
			for warehouse[robot.position.x][robot.position.y+blocksToMove] > 0 {
				blocksToMove++
			}
			if warehouse[robot.position.x][robot.position.y+blocksToMove] == 0 {
				for j := robot.position.y + blocksToMove; j > robot.position.y; j-- {
					warehouse[robot.position.x][j] = warehouse[robot.position.x][j-1]
				}
				warehouse[robot.position.x][robot.position.y] = 0
				robot.position.y++
			}
		case "<":
			for warehouse[robot.position.x][robot.position.y-blocksToMove] > 0 {
				blocksToMove++
			}
			if warehouse[robot.position.x][robot.position.y-blocksToMove] == 0 {
				for j := robot.position.y - blocksToMove; j < robot.position.y; j++ {
					warehouse[robot.position.x][j] = warehouse[robot.position.x][j+1]
				}
				warehouse[robot.position.x][robot.position.y] = 0
				robot.position.y--
			}
		case "^":
			for warehouse[robot.position.x-blocksToMove][robot.position.y] > 0 {
				blocksToMove++
			}
			if warehouse[robot.position.x-blocksToMove][robot.position.y] == 0 {
				for i := robot.position.x - blocksToMove; i < robot.position.x; i++ {
					warehouse[i][robot.position.y] = warehouse[i+1][robot.position.y]
				}
				warehouse[robot.position.x][robot.position.y] = 0
				robot.position.x--
			}
		case "v":
			for warehouse[robot.position.x+blocksToMove][robot.position.y] > 0 {
				blocksToMove++
			}
			if warehouse[robot.position.x+blocksToMove][robot.position.y] == 0 {
				for i := robot.position.x + blocksToMove; i > robot.position.x; i-- {
					warehouse[i][robot.position.y] = warehouse[i-1][robot.position.y]
				}
				warehouse[robot.position.x][robot.position.y] = 0
				robot.position.x++
			}
		}
		// PrettyPrint(warehouse)
	}
	return GetBoxGPSCoordinates(warehouse)
}

func Part2(input []string) int {
	return 1
}

func GetBoxGPSCoordinates(warehouse [][]int) int {
	GPSCoordinates := 0
	for i := range warehouse {
		for j := range warehouse[0] {
			if warehouse[i][j] == 1 {
				GPSCoordinates += 100 * i + j
			}
		}
	}
	return GPSCoordinates
}

func GetInitialWarehouse(input []string) ([][]int, robot) {
	var robot robot
	warehouse := make([][]int, len(input))
	for i, line := range input {
		for j, char := range line {
			if char == '#' {
				warehouse[i] = append(warehouse[i], -1)
			} else if char == '.' {
				warehouse[i] = append(warehouse[i], 0)
			} else if char == 'O' {
				warehouse[i] = append(warehouse[i], 1)
			} else if char == '@' {
				warehouse[i] = append(warehouse[i], 2)
				robot.position = coordinates{i, j}
			}
		}
	}
	return warehouse, robot
}

func GetMovements(input []string) []string {
	var movements []string
	for _, line := range input {
		for _, char := range line {
			movements = append(movements, string(char))
		}
	}
	return movements
}

func PrettyPrint(warehouse [][]int) {
	for i := range warehouse {
		for j := range warehouse[0] {
			if warehouse[i][j] == -1 {
				fmt.Printf("#")
			} else if warehouse[i][j] == 0 {
				fmt.Printf(".")
			} else if warehouse[i][j] == 1 {
				fmt.Printf("0")
			} else {
				fmt.Printf("@")
			}
		}
		fmt.Println("")
	}
}
