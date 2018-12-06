package impl

import (
	"strings"
)

func GetChecksum(boxes string) int {
	boxIds := strings.Split(boxes, "\n")
	numOf2 := 0
	numOf3 := 0

	for _, id := range boxIds {
		hasTwo, hasThree := doesOccur2_3(id)

		if hasTwo {
			numOf2++
		}

		if hasThree {
			numOf3++
		}
	}

	return numOf2 * numOf3
}

func doesOccur2_3(id string) (bool, bool) {
	hasTwo := false
	hasThree := false

	charMap := make(map[string]int)

	for _, char := range strings.Split(id, "") {
		if _, ok := charMap[char]; ok {
			charMap[char]++
		} else {
			charMap[char] = 1
		}
	}

	for _, val := range charMap {
		if val == 2 {
			hasTwo = true
		} else if val == 3 {
			hasThree = true
		}

		if hasTwo && hasThree {
			return hasTwo, hasThree
		}
	}

	return hasTwo, hasThree
}
