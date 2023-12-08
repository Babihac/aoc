package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

type AdjacentEngineRows struct {
	Top     string
	Bottom  string
	Current string
}

type EngineSchematicRow struct {
	Values []SymbolEntry
}

type SymbolEntry struct {
	Value       int
	SymbolValue string
	StartIndex  int
	EndIndex    int
	IsDigit     bool
}

func processNumberString(currentNumberString string, index int, engineSchematicRow *EngineSchematicRow) {
	if currentNumberString != "" {
		numLen := len(currentNumberString)
		number, err := strconv.Atoi(currentNumberString)
		if err != nil {
			log.Fatal(err)
		}

		engineSchematicRow.Values = append(engineSchematicRow.Values, SymbolEntry{Value: number, StartIndex: index - numLen, EndIndex: index - 1, IsDigit: true})

	}
}

func ProcessRow(line string) *EngineSchematicRow {
	var engineSchematicRow EngineSchematicRow
	r := regexp.MustCompile(`[^\d]`)
	var currentNumberString string
	for index, symbol := range line {
		if symbol == '.' {
			processNumberString(currentNumberString, index, &engineSchematicRow)
			currentNumberString = ""
			continue
		}

		if r.MatchString(string(symbol)) {
			processNumberString(currentNumberString, index, &engineSchematicRow)
			currentNumberString = ""
			if symbol == '*' {
				engineSchematicRow.Values = append(engineSchematicRow.Values, SymbolEntry{StartIndex: index, EndIndex: index, IsDigit: false})
			}
		}
		if unicode.IsDigit(symbol) {
			currentNumberString += string(symbol)
		}
	}

	processNumberString(currentNumberString, len(line), &engineSchematicRow)

	return &engineSchematicRow
}

func checkAdjacentSymbols(startIndex int, endIndex int, row string, hasAdjacentSymbol *bool) {
	if row == "" || *hasAdjacentSymbol {
		return
	}

	start := int(math.Max(0, float64(startIndex-1)))
	end := int(math.Min(float64(len(row)-1), float64(endIndex+1)))

	symbolRegexp := regexp.MustCompile(`[^.\d]`)

	for i := start; i <= end; i++ {
		char := string(row[i])
		if symbolRegexp.MatchString(char) {
			*hasAdjacentSymbol = true
		}
	}

}

func checkCurrentRowAdjacentSymbols(startIndex, endIndex int, row string, hasAdjacentSymbol *bool) {

	if *hasAdjacentSymbol {
		return
	}

	symbolRegexp := regexp.MustCompile(`[^.\d]`)
	if startIndex > 0 {
		char := string(row[startIndex-1])
		if symbolRegexp.MatchString(char) {
			*hasAdjacentSymbol = true
		}
	}

	if endIndex < len(row)-1 {
		char := string(row[endIndex+1])
		if symbolRegexp.MatchString(char) {
			*hasAdjacentSymbol = true
		}
	}
}

func SumRowNumbers(adjacentRows *AdjacentEngineRows, row *EngineSchematicRow) int {
	var res int
	for _, entry := range row.Values {
		if !entry.IsDigit {
			continue
		}
		var hasAdjacentSymbol bool
		checkCurrentRowAdjacentSymbols(entry.StartIndex, entry.EndIndex, adjacentRows.Current, &hasAdjacentSymbol)
		checkAdjacentSymbols(entry.StartIndex, entry.EndIndex, adjacentRows.Top, &hasAdjacentSymbol)
		checkAdjacentSymbols(entry.StartIndex, entry.EndIndex, adjacentRows.Bottom, &hasAdjacentSymbol)

		if hasAdjacentSymbol {
			res += entry.Value
		}
	}
	return res
}

func FindCurrentRowGears(index int, row *EngineSchematicRow, gearIndex int) []int {
	var partNumbers []int

	if index > 0 {
		value := row.Values[index-1]

		if value.IsDigit && value.EndIndex == gearIndex-1 {
			partNumbers = append(partNumbers, value.Value)
		}
	}

	if index < len(row.Values)-1 {
		value := row.Values[index+1]

		if value.IsDigit && value.StartIndex == gearIndex+1 {
			partNumbers = append(partNumbers, value.Value)
		}
	}
	return partNumbers
}

func findAdjacentRowGears(row *EngineSchematicRow, gearIndex int) []int {
	var partNumbers []int
	if row == nil {
		return partNumbers
	}

	for _, value := range row.Values {
		if value.IsDigit {
			start := int(math.Max(0, float64(value.StartIndex-1)))
			end := value.EndIndex + 1

			//fmt.Printf("start is %d gear is %d and end is %d\n", start, gearIndex, end)

			if gearIndex >= start && gearIndex <= end {
				partNumbers = append(partNumbers, value.Value)
			}

		}
	}

	return partNumbers

}

func SumGears(bottom, current, top *EngineSchematicRow) int64 {
	var result int64
	for index, value := range current.Values {
		var partNumbers []int
		if !value.IsDigit {
			partNumbers = append(partNumbers, FindCurrentRowGears(index, current, value.StartIndex)...)
			partNumbers = append(partNumbers, findAdjacentRowGears(top, value.StartIndex)...)
			partNumbers = append(partNumbers, findAdjacentRowGears(bottom, value.StartIndex)...)

			if len(partNumbers) == 2 {
				fmt.Println(partNumbers)
				result += int64(partNumbers[0]) * int64(partNumbers[1])
			}
		}
	}
	return result
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	line := ProcessRow(scanner.Text())

	//var totalSum int
	var bottom, current, top *EngineSchematicRow

	var result int64
	for scanner.Scan() {
		nextLine := scanner.Text()
		bottom = ProcessRow(nextLine)
		current = line

		result += SumGears(bottom, current, top)

		top = current
		line = bottom
	}
	bottom = nil
	current = line
	//result += SumGears(bottom, current, top)

	fmt.Printf("Sum gear is: %d", result)
}

//32168385
