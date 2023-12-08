package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ScratchCard struct {
	WinningNumbers []string
	ChosenNumbers  []string
	cardNumber     int
}

func processLine(line string) *ScratchCard {
	values := strings.Split(line, ":")

	if len(values) < 2 {
		return nil
	}

	regex := regexp.MustCompile(`\d+`)

	match := regex.FindString(values[0])

	cardNumber, err := strconv.Atoi(match)

	if err != nil {
		return nil
	}

	numberGroups := strings.Split(values[1], "|")

	if len(numberGroups) < 2 {
		return nil
	}

	return &ScratchCard{WinningNumbers: strings.Fields(numberGroups[0]), ChosenNumbers: strings.Fields(numberGroups[1]), cardNumber: cardNumber}

}

func (s *ScratchCard) CalculateScore() int64 {
	count := make(map[string]bool, len(s.WinningNumbers))
	var winningNumbers int

	for _, num := range s.WinningNumbers {
		count[num] = true
	}

	for _, num := range s.ChosenNumbers {
		if count[num] {
			winningNumbers++
		}
	}
	if winningNumbers == 0 {
		return 0
	}
	return int64(math.Pow(2, float64(winningNumbers)-1))
}

func (s *ScratchCard) WinningNumbersCount() int {
	count := make(map[string]bool, len(s.WinningNumbers))
	var winningNumbers int

	for _, num := range s.WinningNumbers {
		count[num] = true
	}

	for _, num := range s.ChosenNumbers {
		if count[num] {
			winningNumbers++
		}
	}
	return winningNumbers
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	copiesCount := make(map[int]int)

	var res int

	for scanner.Scan() {
		scratchCard := processLine(scanner.Text())

		res += 1 + copiesCount[scratchCard.cardNumber]

		winningNumbersCount := scratchCard.WinningNumbersCount()

		for i := 1; i <= winningNumbersCount; i++ {

			copiesCount[i+scratchCard.cardNumber] += 1 + copiesCount[scratchCard.cardNumber]
		}

	}

	fmt.Printf("Total number of scratch cards is: %d\n", res)

}

// THIS WAS TOOOO SLOW
// for scanner.Scan() {
// 	scratchCard := processLine(scanner.Text())

// 	for i := 0; i < 1+copiesCount[scratchCard.cardNumber]; i++ {
// 		res++
// 		winningNumber := scratchCard.WinningNumbersCount()

// 		for i := 1; i <= winningNumber; i++ {

// 			copiesCount[i+scratchCard.cardNumber] += 1
// 		}
// 	}

// }
