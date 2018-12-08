package impl

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFile(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	// converts content from bytes to 'string'
	str := string(b)

	return str
}

func DoesContain(infiteIDs []string, col string) bool {
	for _, id := range infiteIDs {
		if id == col {
			return true
		}
	}

	return false
}

func GetInfiniteIDs(grid [][]string) []string {
	infiniteIDs := make(map[string]bool)

	maxY := len(grid)
	maxX := len(grid[0])

	for i := 0; i < maxY; i++ {
		id1 := grid[i][0]
		id2 := grid[i][maxX-1]

		infiniteIDs[id1] = true
		infiniteIDs[id2] = true
	}

	for i := 0; i < maxX; i++ {
		id1 := grid[0][i]
		id2 := grid[maxY-1][i]

		infiniteIDs[id1] = true
		infiniteIDs[id2] = true
	}

	avoidIds := []string{}

	for key, _ := range infiniteIDs {
		avoidIds = append(avoidIds, key)
	}

	return avoidIds
}

func InitGrid(height int, width int) [][]string {
	grid := make([][]string, height+1)

	for row := range grid {
		grid[row] = make([]string, width+1)
	}

	for row := range grid {
		for col := range grid[row] {
			grid[row][col] = "."
		}
	}

	return grid
}

func GetLargestXY(points []Point) (int, int) {

	largestX := 0
	largestY := 0

	for _, point := range points {
		if point.x > largestX {
			largestX = point.x
		}

		if point.y > largestY {
			largestY = point.y
		}
	}

	return largestX, largestY
}

func ParsePoints(points []string) []Point {
	parsedPoints := []Point{}
	ids := []string{
		"a", "b", "c", "d", "e", "f", "g",
		"h", "i", "h", "k", "l", "m", "o", "p",
		"q", "r", "s", "t", "u", "v", "w", "x",
		"y", "z", "A", "B", "C", "D", "E", "F",
		"G", "H", "I", "J", "K", "L", "M", "N",
		"O", "P", "Q", "R", "S", "T", "U", "V",
		"W", "x", "Y", "Z"}

	for index, pair := range points {
		yx := strings.Split(pair, ", ")

		y, _ := strconv.Atoi(yx[1])
		x, _ := strconv.Atoi(yx[0])
		id := ids[index]

		parsedPoints = append(parsedPoints, Point{id, x, y})
	}

	return parsedPoints
}

func ManhattenDist(x1 int, y1 int, x2 int, y2 int) int {
	return intAbs(x1-x2) + intAbs(y1-y2)
}

func intAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Point struct {
	id string
	x  int
	y  int
}
