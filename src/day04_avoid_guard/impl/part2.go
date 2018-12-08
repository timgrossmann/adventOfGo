package impl

import (
	"sort"
	"strings"
	"time"
)

func GetHighestAsleepAtMin(guardHours string) (int, int) {
	guardRecords := strings.Split(guardHours, "\n")
	recordList := []record{}

	for _, guardRecord := range guardRecords {
		mappedRecord := mapToRecord(guardRecord)
		recordList = append(recordList, mappedRecord)
	}

	sort.Sort(byTimestamp(recordList))

	mostAsleepID, mostAsleepMin := getMostAsleepMinPerID(recordList)

	return mostAsleepID, mostAsleepMin
}

func getMostAsleepMinPerID(recordList []record) (int, int) {
	timeAsleepMinsPerID := make(map[int]map[int]int)
	guardOnDuty := -1
	fellAsleepTs := time.Now()

	for _, guardRecord := range recordList {
		if timeAsleepMinsPerID[guardOnDuty] == nil {
			timeAsleepMinsPerID[guardOnDuty] = make(map[int]int)
		}

		if guardRecord.eventID == 0 {
			guardOnDuty = guardRecord.guardID
		}

		if guardRecord.eventID == 1 {
			fellAsleepTs = guardRecord.timestamp
		}

		if guardRecord.eventID == 2 {
			for i := fellAsleepTs.Minute(); i < guardRecord.timestamp.Minute(); i++ {
				timeAsleepMinsPerID[guardOnDuty][i]++
			}
		}
	}

	mostAsleepID := -1
	mostAsleepMin := -1
	mostAsleep := -1

	for guardID, minuteMap := range timeAsleepMinsPerID {
		for minute, timesAsleep := range minuteMap {
			if minuteMap[minute] > mostAsleep {
				mostAsleepID = guardID
				mostAsleepMin = minute
				mostAsleep = timesAsleep
			}
		}
	}

	return mostAsleepID, mostAsleepMin
}
