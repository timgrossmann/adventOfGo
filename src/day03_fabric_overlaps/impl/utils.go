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

func parseClaim(claim string) (string, int, int, int, int) {
	split := strings.Split(claim, " ")

	// remove the colon after the x and y def
	split[2] = strings.Replace(split[2], ":", "", 1)
	xy := strings.Split(split[2], ",")
	wh := strings.Split(split[3], "x")

	id := strings.Replace(split[0], "#", "", 1)
	xs, ys := xy[0], xy[1]
	ws, hs := wh[0], wh[1]

	x, _ := strconv.Atoi(xs)
	y, _ := strconv.Atoi(ys)
	w, _ := strconv.Atoi(ws)
	h, _ := strconv.Atoi(hs)

	return id, x, y, w, h
}

func InitFabric(height int, width int) [][]string {
	fabric := make([][]string, height)

	for row := range fabric {
		fabric[row] = make([]string, width)
	}

	for row := range fabric {
		for col := range fabric[row] {
			fabric[row][col] = "."
		}
	}

	return fabric
}

func AddClaimToFabric(fabric [][]string, claim string) (string, int) {
	id, x, y, w, h := parseClaim(claim)

	for row := y; row < y+h; row++ {
		for col := x; col < x+w; col++ {
			if fabric[row][col] != "." {
				fabric[row][col] = "X"
			} else {
				fabric[row][col] = id
			}
		}
	}

	return id, getClaimArea(w, h)
}

func getClaimArea(w int, h int) int {
	return w * h
}
