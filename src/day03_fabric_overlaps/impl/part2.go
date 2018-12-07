package impl

import (
	"strings"
)

func GetIntactClaim(fabricClaims string) string {
	claims := strings.Split(fabricClaims, "\n")
	fabric := InitFabric(1100, 1100)
	areaMapping := make(map[string]int)

	for _, claim := range claims {
		id, area := AddClaimToFabric(fabric, claim)
		areaMapping[id] = area
	}

	intactClaim := findIntactClaim(fabric, areaMapping)

	return intactClaim
}

func findIntactClaim(fabric [][]string, areaMapping map[string]int) string {
	areaCounting := make(map[string]int)

	for row := range fabric {
		for col := range fabric[row] {
			currPos := fabric[row][col]
			areaCounting[currPos]++
		}
	}

	for id, amount := range areaMapping {
		if areaCounting[id] == amount {
			return id
		}
	}

	return ""
}
