package day04

func Part1(input []string) int {
	xmasCount := 0
	for i, line := range input {
		for j, character := range line {
			if character == 'X' {
				if IsXmas(input, i, j, "top") { xmasCount++ }
				if IsXmas(input, i, j, "upper-right") { xmasCount++ }
				if IsXmas(input, i, j, "right") { xmasCount++ }
				if IsXmas(input, i, j, "lower-right") { xmasCount++ }
				if IsXmas(input, i, j, "down") { xmasCount++ }
				if IsXmas(input, i, j, "lower-left") { xmasCount++ }
				if IsXmas(input, i, j, "left") { xmasCount++ }
				if IsXmas(input, i, j, "upper-left") { xmasCount++ }
			}
		}
	}
	return xmasCount
}

func Part2(input []string) int {
	xmasCount := 0
	for i, line := range input {
		for j, character := range line {
			if character == 'A' {
				if IsMasX(input, i, j) { xmasCount++ }
			}
		}
	}
	return xmasCount
}

func IsXmas(input []string, i, j int, direction string) bool {
	var i1, j1, i2, j2, i3, j3 int
	switch direction {
	case "top": i1, j1, i2, j2, i3, j3 = i-1, j, i-2, j, i-3, j
	case "upper-right": i1, j1, i2, j2, i3, j3 = i-1, j+1, i-2, j+2, i-3, j+3
	case "right": i1, j1, i2, j2, i3, j3 = i, j+1, i, j+2, i, j+3
	case "lower-right": i1, j1, i2, j2, i3, j3 = i+1, j+1, i+2, j+2, i+3, j+3
	case "down": i1, j1, i2, j2, i3, j3 = i+1, j, i+2, j, i+3, j
	case "lower-left": i1, j1, i2, j2, i3, j3 = i+1, j-1, i+2, j-2, i+3, j-3
	case "left": i1, j1, i2, j2, i3, j3 = i, j-1, i, j-2, i, j-3
	case "upper-left": i1, j1, i2, j2, i3, j3 = i-1, j-1, i-2, j-2, i-3, j-3
	}
	if i3 < 0 || j3 < 0 || i3 >= len(input) || j3 >= len(input[0]) {
		return false
	}
	return input[i1][j1] == 'M' && input[i2][j2] == 'A' && input[i3][j3] == 'S'
}

func IsMasX(input []string, i, j int) bool {
	if i < 1 || j < 1 || i >= len(input) - 1 || j >= len(input[0]) - 1 {
		return false
	}
	firstDiagonalMas := (input[i-1][j-1] == 'M' && input[i+1][j+1] == 'S') || (input[i-1][j-1] == 'S' && input[i+1][j+1] == 'M')
	secondDiagonalMas := (input[i-1][j+1] == 'M' && input[i+1][j-1] == 'S') || (input[i-1][j+1] == 'S' && input[i+1][j-1] == 'M')
	return firstDiagonalMas && secondDiagonalMas
}
