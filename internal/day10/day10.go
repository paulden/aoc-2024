package day10

import (
	"strconv"
)

type coordinates struct {
	x, y int
}

func (t *coordinates) GetAdjacentTiles(topographicMap [][]int) []coordinates {
	var adjacentTiles []coordinates
	if t.x > 0 {
		adjacentTiles = append(adjacentTiles, coordinates{t.x - 1, t.y})
	}
	if t.y > 0 {
		adjacentTiles = append(adjacentTiles, coordinates{t.x, t.y - 1})
	}
	if t.x < len(topographicMap)-1 {
		adjacentTiles = append(adjacentTiles, coordinates{t.x + 1, t.y})
	}
	if t.y < len(topographicMap[0])-1 {
		adjacentTiles = append(adjacentTiles, coordinates{t.x, t.y + 1})
	}
	return adjacentTiles
}

func (t *coordinates) GetHeight(topographicMap [][]int) int {
	return topographicMap[t.x][t.y]
}

func Part1(input []string) int {
	topographicMap := GetTopographicMap(input)
	hikingTrailsTotalScore := 0
	for i := range topographicMap {
		for j := range topographicMap[0] {
			if topographicMap[i][j] == 0 {
				startTrail := []coordinates{{i, j}}
				hikingTrails := ListHikingTrails(startTrail, topographicMap)
				accessibleNines := ListAccessibleNines(hikingTrails)
				hikingTrailsTotalScore += len(accessibleNines)
			}
		}
	}
	return hikingTrailsTotalScore
}

func Part2(input []string) int {
	topographicMap := GetTopographicMap(input)
	hikingTrailsTotalScore := 0
	for i := range topographicMap {
		for j := range topographicMap[0] {
			if topographicMap[i][j] == 0 {
				startTrail := []coordinates{{i, j}}
				hikingTrails := ListHikingTrails(startTrail, topographicMap)
				hikingTrailsTotalScore += len(hikingTrails)
			}
		}
	}
	return hikingTrailsTotalScore
}

func ListHikingTrails(currentTrail []coordinates, topographicMap [][]int) [][]coordinates {
	currentTile := currentTrail[len(currentTrail)-1]
	tileHeight := currentTile.GetHeight(topographicMap)
	if tileHeight == 9 {
		return [][]coordinates{currentTrail}
	}

	adjacentTiles := currentTile.GetAdjacentTiles(topographicMap)
	hikingTrails := make([][]coordinates, 0)
	for _, tile := range adjacentTiles {
		if tile.GetHeight(topographicMap) == tileHeight+1 {
			nextTrail := make([]coordinates, len(currentTrail))
			copy(nextTrail, currentTrail)
			nextTrail = append(nextTrail, tile)
			hikingTrails = append(hikingTrails, ListHikingTrails(nextTrail, topographicMap)...)
		}
	}

	return hikingTrails
}

func ListAccessibleNines(hikingTrails [][]coordinates) map[coordinates]bool {
	accessibleNines := make(map[coordinates]bool)
	for _, trail := range hikingTrails {
		nineCoordinates := trail[len(trail)-1]
		accessibleNines[nineCoordinates] = true
	}
	return accessibleNines
}

func GetTopographicMap(input []string) [][]int {
	topographicMap := make([][]int, len(input))
	for i, line := range input {
		for _, tile := range line {
			if tile == '.' {
				topographicMap[i] = append(topographicMap[i], -1)
			} else {
				tileHeight, _ := strconv.Atoi(string(tile))
				topographicMap[i] = append(topographicMap[i], tileHeight)
			}
		}
	}
	return topographicMap
}
