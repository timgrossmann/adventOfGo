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

// probably best done with some probablistic structures like a bloom filter
func Contains(list []int64, x int64) bool {
	for _, n := range list {
		if x == n {
			return true
		}
	}
	return false
}
