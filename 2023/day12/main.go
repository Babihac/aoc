package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func processLine(line string, mult int) (string, []int) {
	data := strings.Fields(line)
	recs := data[0]
	var groups []int
	records := strings.Repeat(recs+"?", mult) + recs

	for i := 0; i <= mult; i++ {
		for _, group := range strings.Split(data[1], ",") {
			number, err := strconv.Atoi(group)

			if err != nil {
				log.Fatal(err)
			}

			groups = append(groups, number)
		}
	}

	return records, groups

}

func work(currentString string, template string, groupIndex int, stringIndex int, templateIndex int, damagedSprings int, groups []int) (int, int, int, int, bool) {
	var currentTemplateIndex int
	for i := stringIndex; i < len(currentString); i++ {
		currentTemplateIndex++
		if groupIndex == len(groups) {
			if currentString[i] == '#' {
				return 0, currentTemplateIndex, damagedSprings, groupIndex, false
			} else {
				continue
			}
		}
		var currentGroupSize = groups[groupIndex]
		if currentString[i] == '.' {
			if damagedSprings > 0 && groupIndex == len(groups) {

				return 0, currentTemplateIndex, damagedSprings, groupIndex, false
			}
			if damagedSprings == currentGroupSize {
				groupIndex++

			}
			if damagedSprings > 0 && damagedSprings < currentGroupSize {
				return 0, currentTemplateIndex, damagedSprings, groupIndex, false
			}

			damagedSprings = 0
			continue
		}

		if currentString[i] == '#' {
			damagedSprings++
			if damagedSprings > currentGroupSize {
				return 0, currentTemplateIndex, damagedSprings, groupIndex, false
			}
			if i == len(template)-1 {
				if damagedSprings == groups[groupIndex] {
					//fmt.Printf("increasing group: %d\n", groupIndex)
					groupIndex++
					damagedSprings = 0
				}
			}
		}
	}

	if groupIndex == len(groups) && damagedSprings != 0 {
		return 0, currentTemplateIndex, damagedSprings, groupIndex, false
	}

	if len(currentString) == len(template) && groupIndex == len(groups) {
		return 1, currentTemplateIndex, damagedSprings, groupIndex, true
	} else {
		return 0, currentTemplateIndex, damagedSprings, groupIndex, true
	}
}

func checkRecord(currentString string, template string, groupIndex int, stringIndex int, templateIndex int, damagedSprings int, groups []int) int {

	var argString string
	var result int

	result, currentTemplateIndex, damagedSprings, groupIndex, valid := work(currentString, template, groupIndex, stringIndex, templateIndex, damagedSprings, groups)

	if !valid {
		return 0
	}

	for i := templateIndex; i < len(template); i++ {
		if template[i] != '?' {
			argString += string(template[i])
		} else {
			stringToCheck := currentString + argString + "#"
			result += checkRecord(stringToCheck, template, groupIndex, currentTemplateIndex+stringIndex, i+1, damagedSprings, groups)

			stringToCheck = currentString + argString + "."
			result += checkRecord(stringToCheck, template, groupIndex, currentTemplateIndex+stringIndex, i+1, damagedSprings, groups)
			break
		}
	}
	if argString != "" {
		lastRes, _, _, _, _ := work(currentString+argString, template, groupIndex, currentTemplateIndex+stringIndex, len(template), damagedSprings, groups)
		return result + lastRes

	}
	return result
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var sum float64

	var maxDiff int
	for scanner.Scan() {
		record, groups := processLine(scanner.Text(), 2)
		first := checkRecord("", record, 0, 0, 0, 0, groups)

		record, groups = processLine(scanner.Text(), 3)
		second := checkRecord("", record, 0, 0, 0, 0, groups)

		diff := second / first
		fmt.Printf("mult is %d difference is %d\n", first*first, second-(first*first))
		sum += float64(first * first)
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	fmt.Printf("the result is: %f\n", sum)

	fmt.Println(maxDiff)

}
