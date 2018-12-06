package impl

import (
	"strconv"
	"strings"
)

func GetFrequency(operations string) int64 {
	var resFrequency int64 = 0
	changes := strings.Split(operations, "\n")

	for _, val := range changes {
		intValue, _ := strconv.ParseInt(val, 10, 64)
		resFrequency += intValue
	}

	return resFrequency
}
