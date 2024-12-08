package day08

import "slices"

type coordinates struct {
	x, y int
}

func (c coordinates) IsInBound(input []string) bool {
	return c.x >= 0 && c.x < len(input) && c.y >= 0 && c.y < len(input[0])
}

func Part1(input []string) int {
	antennas := GetAntennasPositions(input)
	antinodes := make(map[coordinates]bool)

	for _, positions := range antennas {
		for i, position := range positions {
			for _, toCompare := range positions[i+1:] {
				antiNode1 := coordinates{position.x + (position.x - toCompare.x), position.y + (position.y - toCompare.y)}
				antiNode2 := coordinates{toCompare.x + (toCompare.x - position.x), toCompare.y + (toCompare.y - position.y)}
				if !slices.Contains(positions, antiNode1) && antiNode1.IsInBound(input) {
					antinodes[antiNode1] = true
				}
				if !slices.Contains(positions, antiNode2) && antiNode2.IsInBound(input) {
					antinodes[antiNode2] = true
				}
			}
		}
	}
	return len(antinodes)
}

func Part2(input []string) int {
	antennas := GetAntennasPositions(input)
	antinodes := make(map[coordinates]bool)

	for _, positions := range antennas {
		for i, position := range positions {
			for _, toCompare := range positions[i+1:] {
				for x, y := position.x, position.y; x >= 0 && x < len(input) && y >= 0 && y < len(input[0]); x, y = x+(position.x-toCompare.x), y+(position.y-toCompare.y) {
					antinodes[coordinates{x, y}] = true
				}
				for x, y := toCompare.x, toCompare.y; x >= 0 && x < len(input) && y >= 0 && y < len(input[0]); x, y = x+(toCompare.x-position.x), y+(toCompare.y-position.y) {
					antinodes[coordinates{x, y}] = true
				}
			}
		}
	}
	return len(antinodes)
}

func GetAntennasPositions(input []string) map[rune][]coordinates {
	antennas := make(map[rune][]coordinates)
	for i, line := range input {
		for j, tile := range line {
			if tile != '.' {
				antennas[tile] = append(antennas[tile], coordinates{i, j})
			}
		}
	}
	return antennas
}
