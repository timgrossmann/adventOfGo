package impl

import (
	"fmt"
	"strings"

	"github.com/arbovm/levenshtein"
)

func CommonCharsOfBoxes(boxes string) string {
	boxIds := strings.Split(boxes, "\n")

	id_1, id_2 := getDistanceOnePair(boxIds)
	commonLetters := getCommonLetters(id_1, id_2)

	return commonLetters
}

// using the levenshtein distance to determine the two strings
// having only one different character
func getDistanceOnePair(boxIds []string) (string, string) {
	for i := 0; i < len(boxIds); i++ {
		id_1 := boxIds[i]

		for j := i; j < len(boxIds); j++ {
			id_2 := boxIds[j]

			if levenshtein.Distance(id_1, id_2) == 1 {
				return id_1, id_2
			}
		}
	}

	fmt.Println("No ids with a distance of 1 found")
	return "", ""
}

func getCommonLetters(id_1 string, id_2 string) string {
	id_1_chars := strings.Split(id_1, "")
	id_2_chars := strings.Split(id_2, "")

	commonChars := []string{}

	for i := 0; i < len(id_1_chars); i++ {
		if id_1_chars[i] == id_2_chars[i] {
			commonChars = append(commonChars, id_1_chars[i])
		}
	}

	return strings.Join(commonChars, "")
}
