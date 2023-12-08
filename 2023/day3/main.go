package main

import (
	"bufio"
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

		if symbol == '*' {
			engineSchematicRow.Values = append(engineSchematicRow.Values, SymbolEntry{StartIndex: index, IsDigit: false})
		}

		if r.MatchString(string(symbol)) {
			processNumberString(currentNumberString, index, &engineSchematicRow)
			currentNumberString = ""
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

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	var adjacentEngineRows AdjacentEngineRows

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	line := scanner.Text()

	var totalSum int

	for scanner.Scan() {
		nextLine := scanner.Text()
		adjacentEngineRows.Bottom = nextLine
		adjacentEngineRows.Current = line
		res := SumRowNumbers(&adjacentEngineRows, ProcessRow(line))

		totalSum += res

		adjacentEngineRows.Top = line
		line = adjacentEngineRows.Bottom
	}

	adjacentEngineRows.Bottom = ""
	adjacentEngineRows.Current = line

	res := SumRowNumbers(&adjacentEngineRows, ProcessRow(line))

	totalSum += res

	log.Printf("Total sum is: %d", totalSum)
}

//498559
