package impl

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadFile(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
	}
	// converts content from bytes to 'string'
	str := string(b)

	return str
}

func myIsUpperCase(char string) bool {
	upperCaseChar := strings.ToUpper(char)

	if char == upperCaseChar {
		return true
	}

	return false
}

func react(cutPolymerChain string) int {
	stack := StringStack{
		values: []string{},
		length: 0,
	}

	for _, char := range cutPolymerChain {
		currChar := string(char)

		if myIsUpperCase(currChar) && myPeek(&stack) == strings.ToLower(currChar) {
			myPop(&stack)
		} else if !myIsUpperCase(currChar) && myPeek(&stack) == strings.ToUpper(currChar) {
			myPop(&stack)
		} else {
			myPush(&stack, currChar)
		}
	}

	return stack.length
}
