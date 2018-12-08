package impl

import (
	"strings"
)

func GetAreaWithTotalDistance(areaPointsList string) int {
	areaPoints := strings.Split(areaPointsList, "\n")

	points := ParsePoints(areaPoints)
	largestX, largestY := GetLargestXY(points)
	grid := InitGrid(largestX+1, largestY+1)
	distanceFieldCount := populateDistanceGrid(grid, points, 10000)

	return distanceFieldCount
}

func populateDistanceGrid(grid [][]string, points []Point, maxDistance int) int {
	distanceCount := 0
	for y, row := range grid {
		for x := range row {
			combinedDistance := 0
			for _, point := range points {
				distance := ManhattenDist(x, y, point.x, point.y)
				combinedDistance += distance
			}

			if combinedDistance < maxDistance {
				grid[y][x] = "#"
				distanceCount++
			}
		}
	}

	return distanceCount
}
