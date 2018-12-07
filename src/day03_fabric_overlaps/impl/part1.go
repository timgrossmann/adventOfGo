package impl

import (
	"strings"
)

func GetOverlaps(fabricClaims string) int {
	claims := strings.Split(fabricClaims, "\n")
	fabric := InitFabric(1100, 1100)

	for _, claim := range claims {
		AddClaimToFabric(fabric, claim)
	}

	amountOfOverlaps := countOverlaps(fabric)

	return amountOfOverlaps
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
