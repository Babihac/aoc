package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func bubbleUpStones(level []rune) int {
	var highestAvailableIndex = 0
	var levelScore int

	for i, placement := range level {
		if placement == '#' {
			highestAvailableIndex = i + 1
		}

		if placement == '.' {
			continue
		}

		if placement == 'O' {

			if i == highestAvailableIndex {
				highestAvailableIndex++
				levelScore += len(level) - i
				continue
			}

			//fmt.Printf("buble up stone on index %d to %d. stone has score %d now\n", i, highestAvailableIndex, len(level)-highestAvailableIndex)
			levelScore += len(level) - highestAvailableIndex
			highestAvailableIndex++
		}
	}

	return levelScore

}

func main() {
	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(file), "\n")

	matrix := make([][]rune, len(data[0]))

	for i := range matrix {
		matrix[i] = make([]rune, len(data))
	}

	for i, line := range data {
		for j, char := range line {
			matrix[j][i] = char
		}
	}

	var result int

	for _, level := range matrix {
		result += bubbleUpStones(level)
	}

	fmt.Printf("Score is %d\n", result)

}
