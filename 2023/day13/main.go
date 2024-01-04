package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func findReflection(matrix []int) int {
	var currentIndex = 1

	for {

		var rowsToCheck = int(math.Min(float64(currentIndex), float64(len(matrix)-currentIndex)))
		var rowsChecked = 0

		if currentIndex == len(matrix) {
			break
		}

		var leftSide = currentIndex - 1
		var rightSide = int(math.Min(float64(currentIndex), float64(len(matrix)-1)))

		for i := 0; i < rowsToCheck; i++ {

			if matrix[leftSide]^matrix[rightSide] == 0 {
				rowsChecked++
			} else {
				break
			}
			leftSide--
			rightSide++
		}

		if rowsChecked == rowsToCheck {
			return currentIndex
		}

		currentIndex++

	}

	return -1

}

func main() {

	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	matrixes := strings.Split(string(file), "\n\n")

	var totalLeftCols int
	var totalAboveRows int

	for _, matrix := range matrixes {
		patterns := strings.Split(matrix, "\n")

		patternMatrix := make([]int, len(patterns))

		transposePatternMatrix := make([]int, len(patterns[0]))

		for i, line := range patterns {
			var binString string

			for j := range line {

				if patterns[i][j] == '.' {
					binString += "0"
				} else if patterns[i][j] == '#' {
					binString += "1"
				}
			}

			number, err := strconv.ParseInt(binString, 2, 64)

			if err != nil {
				log.Fatal(err)
			}
			patternMatrix[i] = int(number)
		}

		for i := range patterns[0] {
			var binString string

			for j := range patterns {

				if patterns[j][i] == '.' {
					binString += "0"
				} else if patterns[j][i] == '#' {
					binString += "1"
				}
			}

			number, err := strconv.ParseInt(binString, 2, 64)

			if err != nil {
				log.Fatal(err)
			}
			transposePatternMatrix[i] = int(number)
		}

		columnsReflection := findReflection(transposePatternMatrix)
		rowsReflection := findReflection(patternMatrix)

		if columnsReflection > 0 {
			totalLeftCols += columnsReflection
			continue
		}

		if rowsReflection > 0 {
			totalAboveRows += rowsReflection
		}

	}

	// for i := range patternMatrix {
	// 	fmt.Printf("%b\n", patternMatrix[i])
	// }

	// fmt.Println()

	// for i := range transposePatternMatrix {
	// 	fmt.Printf("%b\n", transposePatternMatrix[i])
	// }

	fmt.Printf("the result is: %d\n", totalLeftCols+(100*totalAboveRows))

}
