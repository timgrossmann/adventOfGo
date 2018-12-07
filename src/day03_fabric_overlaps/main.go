package main

import (
	"day03_fabric_overlaps/impl"
	"fmt"
)

func main() {
	fileContent := impl.ReadFile("resources/fabric_claims.txt")

	amountOfOverlaps := impl.GetOverlaps(fileContent)
	intactClaim := impl.GetIntactClaim(fileContent)

	fmt.Println(amountOfOverlaps, intactClaim)
}
