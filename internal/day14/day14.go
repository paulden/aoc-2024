package day14

import (
	"fmt"
	"regexp"
	"strconv"
)

type coordinates struct {
	x, y int
}

type robot struct {
	position, vector coordinates
}

func (r *robot) Move(h, w int) {
	nextX := (r.position.x + r.vector.x + w) % w
	nextY := (r.position.y + r.vector.y + h) % h
	r.position = coordinates{nextX, nextY}
}

func Part1(input []string, h, w int) int {
	robotsByPosition := GetRobotsByPosition(input)
	for range 100 {
		nextRobotsByPosition := make(map[coordinates][]robot)
		for _, robots := range robotsByPosition {
			for _, robot := range robots {
				robot.Move(h, w)
				nextRobotsByPosition[robot.position] = append(nextRobotsByPosition[robot.position], robot)
			}
		}
		robotsByPosition = nextRobotsByPosition
	}

	mh := h / 2
	mw := w / 2

	first := CountRobotsInQuadrant(0, 0, mw, mh, robotsByPosition)
	second := CountRobotsInQuadrant(mw+1, 0, w, mh, robotsByPosition)
	third := CountRobotsInQuadrant(0, mh+1, mw, h, robotsByPosition)
	fourth := CountRobotsInQuadrant(mw+1, mh+1, w, h, robotsByPosition)

	return first * second * third * fourth
}

func Part2(input []string) int {
	robotsByPosition := GetRobotsByPosition(input)
	secondsElapsed := 0
	for !IsChristmasTree(robotsByPosition) {
		secondsElapsed++
		nextRobotsByPosition := make(map[coordinates][]robot)
		for _, robots := range robotsByPosition {
			for _, robot := range robots {
				robot.Move(103, 101)
				nextRobotsByPosition[robot.position] = append(nextRobotsByPosition[robot.position], robot)
			}
		}
		robotsByPosition = nextRobotsByPosition
	}
	PrettyPrint(robotsByPosition)
	return secondsElapsed
}

func GetRobotsByPosition(input []string) map[coordinates][]robot {
	robotsByPosition := make(map[coordinates][]robot)
	for _, line := range input {
		robot := GetRobot(line)
		robotsByPosition[robot.position] = append(robotsByPosition[robot.position], robot)
	}
	return robotsByPosition
}

func GetRobot(line string) robot {
	r := regexp.MustCompile(`(-?\d+)`)
	numbers := r.FindAllString(line, -1)
	px, _ := strconv.Atoi(numbers[0])
	py, _ := strconv.Atoi(numbers[1])
	vx, _ := strconv.Atoi(numbers[2])
	vy, _ := strconv.Atoi(numbers[3])
	return robot{coordinates{px, py}, coordinates{vx, vy}}
}

func CountRobotsInQuadrant(startX, startY, endX, endY int, robotsByPosition map[coordinates][]robot) int {
	robotsCount := 0
	for i := startX; i < endX; i++ {
		for j := startY; j < endY; j++ {
			robotsCount += len(robotsByPosition[coordinates{i, j}])
		}
	}
	return robotsCount
}

func IsChristmasTree(robotsByPosition map[coordinates][]robot) bool {
	for i := range 101 {
		for j := range 103 {
			// Fifteen consecutive robots should be enough to say it's a Christmas tree, right? Right?
			if len(robotsByPosition[coordinates{i, j}]) > 0 && HasConsecutiveRobots(i, j, 15, robotsByPosition) {
				return true
			}
		}
	}
	return false
}

func HasConsecutiveRobots(x, y, size int, robotsByPosition map[coordinates][]robot) bool {
	for i := x; i < x+size; i++ {
		if len(robotsByPosition[coordinates{i, y}]) == 0 {
			return false
		}
	}
	return true
}
func PrettyPrint(robotsByPosition map[coordinates][]robot) {
	for j := range 103 {
		for i := range 101 {
			if len(robotsByPosition[coordinates{i, j}]) > 0 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Println("")
	}
}
