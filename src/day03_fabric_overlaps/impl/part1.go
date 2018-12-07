package impl

import (
	"strconv"
	"strings"
)

func GetOverlaps(fabricClaims string) int {
	claims := strings.Split(fabricClaims, "\n")
	fabric := initFabric(1100, 1100)

	for _, claim := range claims {
		addClaimToFabric(fabric, claim)
	}

	amountOfOverlaps := countOverlaps(fabric)

	return amountOfOverlaps
}

func addClaimToFabric(fabric [][]string, claim string) {
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
}

func parseClaim(claim string) (string, int, int, int, int) {
	split := strings.Split(claim, " ")

	// remove the colon after the x and y def
	split[2] = strings.Replace(split[2], ":", "", 1)
	xy := strings.Split(split[2], ",")
	wh := strings.Split(split[3], "x")

	id := strings.Replace(split[0], "#", "", 1)
	xs, ys := xy[1], xy[0]
	ws, hs := wh[0], wh[1]

	x, _ := strconv.Atoi(xs)
	y, _ := strconv.Atoi(ys)
	w, _ := strconv.Atoi(ws)
	h, _ := strconv.Atoi(hs)

	return id, x, y, w, h
}

func initFabric(height int, width int) [][]string {
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

func countOverlaps(fabric [][]string) int {

	amountOfOverlaps := 0

	for row := range fabric {
		for col := range fabric[row] {
			if fabric[row][col] == "X" {
				amountOfOverlaps++
			}
		}
	}

	return amountOfOverlaps
}
