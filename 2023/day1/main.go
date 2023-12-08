package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

//53623
//54815

func convertFirstDigitWord(line string) string {
	digits := []string{"one.1", "two.2", "three.3", "four.4", "five.5", "six.6", "seven.7", "eight.8", "nine.9"}
	var firstDigitWord string
	var digitToReplace string
	var lowestIndex = math.MaxInt
	for _, digit := range digits {
		wordAndDigit := strings.Split(digit, ".")
		index := strings.Index(line, wordAndDigit[0])

		if index == -1 {
			continue
		}
		if index < lowestIndex {
			lowestIndex = index
			firstDigitWord = wordAndDigit[0]
			digitToReplace = wordAndDigit[1]

		}
	}
	line = strings.Replace(line, firstDigitWord, digitToReplace, 1)

	return line
}

func convertLastDigitWord(line string) string {
	digits := []string{"one.1", "two.2", "three.3", "four.4", "five.5", "six.6", "seven.7", "eight.8", "nine.9"}
	var firstDigitWord string
	var digitToReplace string
	var highestIndex = -1
	for _, digit := range digits {
		wordAndDigit := strings.Split(digit, ".")
		index := strings.LastIndex(line, wordAndDigit[0])

		if index == -1 {
			continue
		}

		if index > highestIndex {
			highestIndex = index
			firstDigitWord = wordAndDigit[0]
			digitToReplace = wordAndDigit[1]

		}
	}

	line = strings.Replace(line, firstDigitWord, digitToReplace, 1)

	return line
}

func processLine(line string) int {
	line = convertLastDigitWord(convertFirstDigitWord(line))
	var firstDigit = '0'
	var lastDigit = '0'
	var firstAdded = false
	for _, x := range line {
		if unicode.IsDigit(x) {
			lastDigit = x

			if !firstAdded {
				firstAdded = true
				firstDigit = x
			}
		}
	}

	combined := string(firstDigit) + string(lastDigit)
	result, err := strconv.Atoi(combined)

	if err != nil {
		log.Fatal(err)
	}

	return result

}

func main() {
	file, err := os.Open("./input.txt")
	var result int
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		result += processLine(scanner.Text())
	}

	fmt.Printf("Result is: %d\n", result)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}
