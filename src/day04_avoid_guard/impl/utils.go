package impl

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"time"
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

func mapToRecord(txtRecord string) record {
	timeFormat := "2006-01-02 15:04"
	timestamp, event := splitText(txtRecord)
	parsedTimestamp, _ := time.Parse(timeFormat, timestamp)

	eventID := 0
	guardID := -1

	switch event {
	case "falls asleep":
		eventID = 1
		break
	case "wakes up":
		eventID = 2
		break
	default:
		guardID = extractID(event)
		break
	}

	return record{
		timestamp: parsedTimestamp,
		guardID:   guardID,
		eventID:   eventID,
	}
}

func splitText(txtRecord string) (string, string) {
	splitRegex := regexp.MustCompile(`\[(?P<timestamp>.*)\] (?P<event>.*)`)

	timestamp := splitRegex.FindStringSubmatch(txtRecord)[1]
	event := splitRegex.FindStringSubmatch(txtRecord)[2]

	return timestamp, event
}

func extractID(event string) int {
	regex := regexp.MustCompile(`.* #(?P<id>\d*) .*`)

	guardID, _ := strconv.Atoi(regex.FindStringSubmatch(event)[1])

	return guardID
}
