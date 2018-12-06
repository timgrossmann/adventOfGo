package impl

import (
	"strconv"
	"strings"
)

func GetRepeatingFrequency(operations string) int64 {
	changes := strings.Split(operations, "\n")
	seenFrequencies := []int64{0}
	var resFrequency int64 = 0

	parsedValues := []int64{}
	for _, val := range changes {
		intValue, _ := strconv.ParseInt(val, 10, 64)
		parsedValues = append(parsedValues, intValue)
	}

	for {
		for _, val := range parsedValues {
			resFrequency += val

			if Contains(seenFrequencies, resFrequency) {
				return resFrequency
			}

			seenFrequencies = append(seenFrequencies, resFrequency)
		}
	}
}
