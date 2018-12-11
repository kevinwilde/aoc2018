package main

import (
	"aoc2018/utils"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strconv"
	"time"
)

func main() {
	input := utils.GetInput("input.txt")
	events := groupEventsByGuard(sortEventsByTime(convertInputToEvents(input)))
	guardsToMinToNumTimesAsleep := mapGuardToMinToNumTimesAsleep(events)
	fmt.Println(part1(guardsToMinToNumTimesAsleep))
	fmt.Println(part2(guardsToMinToNumTimesAsleep))
}

func part1(guardsToMinToNumTimesAsleep map[int]map[int]int) int {
	sleepiestGuardID := findKeyWithMaxValue(mapGuardToMinsAsleep(guardsToMinToNumTimesAsleep))
	sleepiestMinute := findKeyWithMaxValue(guardsToMinToNumTimesAsleep[sleepiestGuardID])
	return sleepiestGuardID * sleepiestMinute
}

func part2(guardsToMinToNumTimesAsleep map[int]map[int]int) int {
	maxSleepOccurrences := 0
	sleepiestGuardID, sleepiestMinute := -1, -1

	for guardID, minToNumTimesAsleep := range guardsToMinToNumTimesAsleep {
		m := findKeyWithMaxValue(minToNumTimesAsleep)
		sleepOccurrences := minToNumTimesAsleep[m]
		if sleepOccurrences > maxSleepOccurrences {
			maxSleepOccurrences = sleepOccurrences
			sleepiestGuardID = guardID
			sleepiestMinute = m
		}
	}

	return sleepiestGuardID * sleepiestMinute
}

type eventType int

const (
	startShift eventType = 0
	fallAsleep eventType = 1
	wakeUp     eventType = 2
)

type event struct {
	datetime time.Time
	action   eventType
	guardID  int
}

func convertInputToEvents(input []string) []*event {
	var events []*event

	reg := regexp.MustCompile(`^\[(\d\d\d\d-\d\d-\d\d \d\d:\d\d)\] (wakes up|falls asleep|Guard #(\d+) begins shift)$`)

	for _, s := range input {
		match := reg.FindStringSubmatch(s)
		if len(match) < 1 {
			log.Fatal("Input did not match regex")
		}

		eventTime, err := time.Parse("2006-01-02 15:04", match[1])
		if err != nil {
			log.Fatal(err)
		}

		var ev event
		if match[3] != "" {
			guardID, err := strconv.Atoi(match[3])
			if err != nil {
				log.Fatal(err)
			}
			ev = event{eventTime, startShift, guardID}
		} else if match[2] == "wakes up" {
			ev = event{eventTime, wakeUp, -1}
		} else if match[2] == "falls asleep" {
			ev = event{eventTime, fallAsleep, -1}
		}
		events = append(events, &ev)
	}

	return events
}

func mapGuardToMinToNumTimesAsleep(events map[int][]*event) map[int]map[int]int {
	guardsToMinToNumTimesAsleep := make(map[int]map[int]int)
	for guardID, events := range events {
		for i, ev := range events {
			if ev.action == fallAsleep {
				for m := ev.datetime.Minute(); m < events[i+1].datetime.Minute(); m++ {
					if guardsToMinToNumTimesAsleep[guardID] == nil {
						guardsToMinToNumTimesAsleep[guardID] = make(map[int]int)
					}
					guardsToMinToNumTimesAsleep[guardID][m]++
				}
			}
		}
	}
	return guardsToMinToNumTimesAsleep
}

func sortEventsByTime(events []*event) []*event {
	sort.Slice(events, func(i, j int) bool {
		return events[i].datetime.Before(events[j].datetime)
	})
	return events
}

func groupEventsByGuard(events []*event) map[int][]*event {
	groupedEvents := make(map[int][]*event)

	lastGuardID := -1
	for _, ev := range events {
		switch ev.action {
		case startShift:
			if groupedEvents[ev.guardID] == nil {
				groupedEvents[ev.guardID] = append(groupedEvents[ev.guardID], ev)
			}
			lastGuardID = ev.guardID
		case fallAsleep:
			fallthrough
		case wakeUp:
			ev.guardID = lastGuardID
			groupedEvents[ev.guardID] = append(groupedEvents[ev.guardID], ev)
		}
	}
	return groupedEvents
}

func mapGuardToMinsAsleep(guardsToMinToNumTimesAsleep map[int]map[int]int) map[int]int {
	minsAsleep := make(map[int]int)
	for guardID, minsToNumTimesAsleep := range guardsToMinToNumTimesAsleep {
		mins := 0
		for _, times := range minsToNumTimesAsleep {
			mins += times
		}
		minsAsleep[guardID] = mins
	}
	return minsAsleep
}

func findKeyWithMaxValue(kvMap map[int]int) int {
	curMaxVal, curMaxKey := -1, -1
	for k, v := range kvMap {
		if v > curMaxVal {
			curMaxVal = v
			curMaxKey = k
		}
	}
	return curMaxKey
}
