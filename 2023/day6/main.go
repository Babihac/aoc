package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type BoatRecord struct {
	time     int
	distance int
}

type RaceTable struct {
	records []BoatRecord
}

func NewRecordsTable(times []string, distances []string) *RaceTable {
	if len(distances) != len(times) {
		return nil
	}

	var table RaceTable

	for i, distance := range distances {
		numericTime, err := strconv.Atoi(times[i])

		if err != nil {
			return nil
		}

		numericDistance, err := strconv.Atoi(distance)

		if err != nil {
			return nil
		}

		table.records = append(table.records, BoatRecord{time: numericTime, distance: numericDistance})
	}

	return &table
}

func calculatePossibleDistances(time int, limitDistance int) int {
	var minVelocity = 0
	var maxVelocity = time

	for {
		if limitDistance >= ((time - minVelocity) * minVelocity) {
			minVelocity++
			continue
		} else if limitDistance >= ((time - maxVelocity) * maxVelocity) {
			maxVelocity--
			continue
		}
		break
	}
	return maxVelocity - minVelocity + 1
}

func main() {
	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(file), "\n")

	times := strings.Fields(data[0])
	distances := strings.Fields(data[1])

	table := NewRecordsTable(times[1:], distances[1:])

	var res = 1

	for _, t := range table.records {
		res *= calculatePossibleDistances(t.time, t.distance)
	}
	fmt.Printf("result is; %d'\n", res)
}
