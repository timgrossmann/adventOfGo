package main

import (
	"day02_box_ids/impl"
	"fmt"
)

func main() {
	fileContent := impl.ReadFile("resources/box_ids.txt")

	checksum := impl.GetChecksum(fileContent)
	commonLetters := impl.CommonCharsOfBoxes(fileContent)

	fmt.Println(checksum, commonLetters)
}
