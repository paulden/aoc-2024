package day12

type coordinates struct {
	x, y int
}

func (c *coordinates) GetAdjacentCoordinate(input []string) []coordinates {
	var adjacentCoordinates []coordinates
	if c.x > 0 {
		adjacentCoordinates = append(adjacentCoordinates, coordinates{c.x - 1, c.y})
	}
	if c.y > 0 {
		adjacentCoordinates = append(adjacentCoordinates, coordinates{c.x, c.y - 1})
	}
	if c.x < len(input)-1 {
		adjacentCoordinates = append(adjacentCoordinates, coordinates{c.x + 1, c.y})
	}
	if c.y < len(input[0])-1 {
		adjacentCoordinates = append(adjacentCoordinates, coordinates{c.x, c.y + 1})
	}
	return adjacentCoordinates
}

type plot struct {
	plantType string
	plants    map[coordinates]bool
}

func (p *plot) ListAllPlants(startTile coordinates, input []string, visitedTiles map[coordinates]bool) {
	for plant := range p.plants {
		adjacentPlants := plant.GetAdjacentCoordinate(input)
		for _, adjacentPlant := range adjacentPlants {
			if !visitedTiles[adjacentPlant] {
				if string(input[adjacentPlant.x][adjacentPlant.y]) == p.plantType {
					visitedTiles[adjacentPlant] = true
					p.plants[adjacentPlant] = true
					p.ListAllPlants(adjacentPlant, input, visitedTiles)
				}
			}
		}
	}
}

func (p *plot) GetArea() int {
	return len(p.plants)
}
func (p *plot) GetPerimeter() int {
	perimeter := 0
	for plant := range p.plants {
		baseFenceLength := 4
		if p.plants[coordinates{plant.x, plant.y - 1}] {
			baseFenceLength--
		}
		if p.plants[coordinates{plant.x, plant.y + 1}] {
			baseFenceLength--
		}
		if p.plants[coordinates{plant.x - 1, plant.y}] {
			baseFenceLength--
		}
		if p.plants[coordinates{plant.x + 1, plant.y}] {
			baseFenceLength--
		}
		perimeter += baseFenceLength
	}
	return perimeter
}

func (p *plot) GetSides() int {
	sides := 0
	for plant := range p.plants {
		topPlant := p.plants[coordinates{plant.x - 1, plant.y}]
		topLeftPlant := p.plants[coordinates{plant.x - 1, plant.y - 1}]
		leftPlant := p.plants[coordinates{plant.x, plant.y - 1}]
		bottomLeftPlant := p.plants[coordinates{plant.x + 1, plant.y - 1}]
		bottomPlant := p.plants[coordinates{plant.x + 1, plant.y}]
		bottomRightPlant := p.plants[coordinates{plant.x + 1, plant.y + 1}]
		rightPlant := p.plants[coordinates{plant.x, plant.y + 1}]
		topRightPlant := p.plants[coordinates{plant.x - 1, plant.y + 1}]

		if topPlant && leftPlant && !topLeftPlant || !topPlant && !leftPlant {
			sides++
		}
		if topPlant && rightPlant && !topRightPlant || !topPlant && !rightPlant {
			sides++
		}
		if bottomPlant && leftPlant && !bottomLeftPlant || !bottomPlant && !leftPlant {
			sides++
		}
		if bottomPlant && rightPlant && !bottomRightPlant || !bottomPlant && !rightPlant {
			sides++
		}

	}
	return int(sides)
}

func Part1(input []string) int {
	plots := GetPlots(input)
	fencePrice := 0
	for _, plot := range plots {
		fencePrice += plot.GetArea() * plot.GetPerimeter()
	}
	return fencePrice
}

func Part2(input []string) int {
	plots := GetPlots(input)
	fencePrice := 0
	for _, plot := range plots {
		fencePrice += plot.GetArea() * plot.GetSides()
	}
	return fencePrice
}

func GetPlots(input []string) []plot {
	var plots []plot
	visitedTiles := make(map[coordinates]bool)
	for i, line := range input {
		for j, char := range line {
			if visitedTiles[coordinates{i, j}] {
				continue
			} else {
				visitedTiles[coordinates{i, j}] = true
				newPlot := plot{string(char), map[coordinates]bool{{i, j}: true}}
				newPlot.ListAllPlants(coordinates{i, j}, input, visitedTiles)
				plots = append(plots, newPlot)
			}
		}
	}
	return plots
}

func ParsePlot(currentTile coordinates, input []string) plot {
	plot := plot{string(input[currentTile.x][currentTile.y]), map[coordinates]bool{currentTile: true}}
	return plot
}
