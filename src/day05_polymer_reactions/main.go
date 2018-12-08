package main

import (
	"day05_polymer_reactions/impl"
	"fmt"
)

func main() {
	fileContent := impl.ReadFile("resources/polymer_chains.txt")

	remainingPolymerLen := impl.GetRemainingPolymer(fileContent)
	shortestPolymerLen := impl.GetShortestPolymerLen(fileContent)

	fmt.Println(remainingPolymerLen, shortestPolymerLen)
}
