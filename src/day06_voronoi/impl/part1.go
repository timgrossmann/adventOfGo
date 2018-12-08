package impl

import (
	"math"
	"strings"
)

func GetLargestArea(areaPointsList string) int {
	areaPoints := strings.Split(areaPointsList, "\n")

	points := ParsePoints(areaPoints)
	largestX, largestY := GetLargestXY(points)
	grid := InitGrid(largestX+1, largestY+1)
	populateGrid(grid, points)

	largestArea := getLargestArea(grid)

	return largestArea
}

func populateGrid(grid [][]string, points []Point) {
	for y, row := range grid {
		for x := range row {
			shortestDistance := math.MaxInt32
			occupator := "."

			for _, point := range points {
				distance := ManhattenDist(x, y, point.x, point.y)

				if distance < shortestDistance {
					shortestDistance = distance
					occupator = point.id
				} else if distance == shortestDistance {
					occupator = "."
				}
			}

			grid[y][x] = occupator
		}
	}
}

func getLargestArea(grid [][]string) int {
	areaPerPoint := make(map[string]int)
	infiteIDs := GetInfiniteIDs(grid)

	for _, row := range grid {
		for _, col := range row {
			if col == "." {
				continue
			}

			areaPerPoint[col]++
		}
	}

	largestArea := 0

	for id, amount := range areaPerPoint {
		if amount > largestArea && !DoesContain(infiteIDs, id) {
			largestArea = amount
		}
	}

	return largestArea
}
