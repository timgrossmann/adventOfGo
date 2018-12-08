package main

import (
	"day04_avoid_guard/impl"
	"fmt"
)

func main() {
	fileContent := impl.ReadFile("resources/guard_hours.txt")

	guardID, guardHour := impl.GetGuardAndHour(fileContent)
	guardID2, guardHour2 := impl.GetHighestAsleepAtMin(fileContent)

	fmt.Println(guardID*guardHour, guardID2*guardHour2)
}
