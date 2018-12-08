package impl

import (
	"math"
	"strings"
)

func GetShortestPolymerLen(polymerChain string) int {
	chars := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i",
		"j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z"}

	mostReactedLen := math.MaxInt32

	for _, char := range chars {
		upperChar := strings.ToUpper(char)

		charsRemoved := strings.Replace(polymerChain, char, "", -1)
		charsRemoved = strings.Replace(charsRemoved, upperChar, "", -1)

		reactedLen := react(charsRemoved)

		if reactedLen < mostReactedLen {
			mostReactedLen = reactedLen
		}
	}

	return mostReactedLen
}
