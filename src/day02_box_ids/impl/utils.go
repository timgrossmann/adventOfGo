package impl

import (
	"fmt"
	"io/ioutil"
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
