package impl

import (
	"sort"
	"strings"
	"time"
)

func GetGuardAndHour(guardHours string) (int, int) {
	guardRecords := strings.Split(guardHours, "\n")
	recordList := []record{}

	for _, guardRecord := range guardRecords {
		mappedRecord := mapToRecord(guardRecord)
		recordList = append(recordList, mappedRecord)
	}

	sort.Sort(byTimestamp(recordList))

	mostAsleepID := getMostMinAsleepID(recordList)
	mostAsleepMin := getMostAsleepMin(recordList, mostAsleepID)

	return mostAsleepID, mostAsleepMin
}

func getMostAsleepMin(recordList []record, mostAsleepID int) int {
	timeAsleepMins := make(map[int]int)
	guardOnDuty := -1
	fellAsleepTs := time.Now()

	for _, guardRecord := range recordList {
		if guardRecord.eventID == 0 {
			guardOnDuty = guardRecord.guardID
		}

		if guardOnDuty != mostAsleepID {
			continue
		}

		if guardRecord.eventID == 1 {
			fellAsleepTs = guardRecord.timestamp
		}

		if guardRecord.eventID == 2 {
			for i := fellAsleepTs.Minute(); i < guardRecord.timestamp.Minute(); i++ {
				timeAsleepMins[i]++
			}
		}
	}

	mostAsleepMin := -1
	mostAsleep := -1

	for minute, timesAsleep := range timeAsleepMins {
		if timeAsleepMins[minute] > mostAsleep {
			mostAsleepMin = minute
			mostAsleep = timesAsleep
		}
	}

	return mostAsleepMin
}

func getMostMinAsleepID(recordList []record) int {
	timeAsleepIds := make(map[int]int)
	guardOnDutyID := -1
	fellAsleepTs := time.Now()

	for _, guardRecord := range recordList {
		if guardRecord.eventID == 0 {
			guardOnDutyID = guardRecord.guardID
		}

		if guardRecord.eventID == 1 {
			fellAsleepTs = guardRecord.timestamp
		}

		if guardRecord.eventID == 2 {
			timeDifference := guardRecord.timestamp.Sub(fellAsleepTs)
			minDiff := int(timeDifference.Minutes())
			hourDiff := int(timeDifference.Hours())
			timeAsleepIds[guardOnDutyID] += minDiff + hourDiff*60
		}
	}

	mostAsleepID := -1
	longestAsleep := -1

	for guardID, timeAsleep := range timeAsleepIds {
		if timeAsleepIds[guardID] > longestAsleep {
			mostAsleepID = guardID
			longestAsleep = timeAsleep
		}
	}

	return mostAsleepID
}
