package main

import (
	"day01_frequency/impl"
	"fmt"
)

func main() {
	fileContent := impl.ReadFile("resources/frequency_input.txt")

	finalFrequency := impl.GetFrequency(fileContent)
	repeatedFrequency := impl.GetRepeatingFrequency(fileContent)

	fmt.Println("Final Frequency", finalFrequency)
	fmt.Println("Repeated Frequency", repeatedFrequency)
}
