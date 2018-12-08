package main

import (
	"day06_voronoi/impl"
	"fmt"
)

func main() {
	fileContent := impl.ReadFile("resources/locations.txt")

	largestArea := impl.GetLargestArea(fileContent)
	areaWithTotalDistance := impl.GetAreaWithTotalDistance(fileContent)

	fmt.Println(largestArea, areaWithTotalDistance)
}
