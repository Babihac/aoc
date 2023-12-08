package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type SourceToDestMap struct {
	DestStart   int
	SourceStart int
	interval    int
}

func NewSourceToDestMap(dest, source, interval string) *SourceToDestMap {
	destValue, err := strconv.Atoi(dest)

	if err != nil {
		log.Fatal(err)
	}

	sourceValue, err := strconv.Atoi(source)

	if err != nil {
		log.Fatal(err)
	}

	intervalValue, err := strconv.Atoi(interval)

	if err != nil {
		log.Fatal(err)
	}

	return &SourceToDestMap{interval: intervalValue, DestStart: destValue, SourceStart: sourceValue}

}

func mapSeedIntervalToLocation(seedInterval []int, mappings [][]SourceToDestMap) int {
	mappedLocation := seedInterval[0]
	seedsInInterval := seedInterval[1]
	var intervalMin int = math.MaxInt
	var processedSeeds = 0
	var reduction int

	for {
		if processedSeeds >= seedsInInterval {
			break
		}
		reduction = seedsInInterval
		var reductionSet = false
		for _, mapping := range mappings {
			mappingFound := false
			for _, sourceToDest := range mapping {
				if mappingFound {
					break
				}

				if mappedLocation >= sourceToDest.SourceStart && mappedLocation < sourceToDest.SourceStart+sourceToDest.interval {
					mappingFound = true
					mappedLocation = mappedLocation - sourceToDest.SourceStart + sourceToDest.DestStart
					reduction = int(math.Min(float64(reduction), float64(sourceToDest.interval-(mappedLocation-sourceToDest.DestStart))))
					reductionSet = true
				}
			}
		}
		if intervalMin > mappedLocation {
			intervalMin = mappedLocation
		}

		if !reductionSet {
			reduction = 1
		}

		processedSeeds += reduction
		mappedLocation = seedInterval[0] + processedSeeds
	}

	return intervalMin
}

func main() {
	var mapping [][]SourceToDestMap
	var lowestLocacion = math.MaxInt

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seedString := strings.TrimPrefix(scanner.Text(), "seeds: ")
	seedArr := strings.Split(seedString, " ")
	seeds := make([][]int, 0)

	for i := 0; i < len(seedArr); i += 2 {
		start, err := strconv.Atoi(seedArr[i])

		if err != nil {
			log.Fatal(err)
		}

		interval, err := strconv.Atoi(seedArr[i+1])

		if err != nil {
			log.Fatal(err)
		}

		seeds = append(seeds, []int{start, interval})
	}

	for scanner.Scan() {
		input := scanner.Text()

		if input == "" {
			continue
		}

		if strings.ContainsRune(input, ':') {
			mapping = append(mapping, []SourceToDestMap{})
		} else {
			mappingValues := strings.Fields(input)

			if len(mappingValues) < 3 {
				log.Fatal("invalid value")
			}

			sourceToDestMap := NewSourceToDestMap(mappingValues[0], mappingValues[1], mappingValues[2])

			mapping[len(mapping)-1] = append(mapping[len(mapping)-1], *sourceToDestMap)

		}
	}

	for _, seedRange := range seeds {
		currentLocation := mapSeedIntervalToLocation(seedRange, mapping)

		//fmt.Printf("current location: %d for seed: %d\n", currentLocation, firstSeed+i)

		if currentLocation < lowestLocacion {
			lowestLocacion = currentLocation
		}
	}

	fmt.Printf("Lowest location is: %d\n", lowestLocacion)

}
