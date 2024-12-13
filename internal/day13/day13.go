package day13

import (
	"regexp"
	"strconv"
)

type coordinates struct {
	x, y int
}

type machine struct {
	buttonA, buttonB, prize coordinates
}

func (m machine) GetPressesRequiredToWin() (int, int) {
	determinant := m.buttonA.x*m.buttonB.y - m.buttonA.y*m.buttonB.x
	if determinant == 0 {
		return -1, -1
	} else {
		aNumerator := m.prize.x*m.buttonB.y - m.prize.y*m.buttonB.x
		bNumerator := m.prize.y*m.buttonA.x - m.prize.x*m.buttonA.y
		if aNumerator % determinant == 0 && bNumerator % determinant == 0 {
			return aNumerator / determinant, bNumerator / determinant
		}
		return -1, -1
	}
}

func Part1(input []string) int {
	tokensToSpend := 0
	for i := 0; i < len(input); i += 4 {
		machine := machine{GetCoordinates(input[i], 0), GetCoordinates(input[i+1], 0), GetCoordinates(input[i+2], 0)}
		a, b := machine.GetPressesRequiredToWin()
		if a >= 0 && a <= 100 && b >= 0 && b <= 100 {
			tokensToSpend += 3*a + b
		}
	}
	return tokensToSpend
}

func Part2(input []string) int {
	tokensToSpend := 0
	for i := 0; i < len(input); i += 4 {
		machine := machine{GetCoordinates(input[i], 0), GetCoordinates(input[i+1], 0), GetCoordinates(input[i+2], 10000000000000)}
		a, b := machine.GetPressesRequiredToWin()
		if a >= 0 && b >= 0 {
			tokensToSpend += 3*a + b
		}
	}
	return tokensToSpend
}

func GetCoordinates(line string, padding int) coordinates {
	r := regexp.MustCompile(`(\d+)`)
	numbers := r.FindAllString(line, -1)
	x, _ := strconv.Atoi(numbers[0])
	y, _ := strconv.Atoi(numbers[1])
	return coordinates{x+padding, y+padding}
}
