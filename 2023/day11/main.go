package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

type Coord struct {
	X        int
	Y        int
	Distance int
}

func expandUniverse(universe [][]rune) [][]rune {
	var res [][]rune

	var emptyColumns []int
	for i := 0; i < len(universe[0]); i++ {
		var columnWithoutGalaxy = true
		for j := 0; j < len(universe); j++ {
			if universe[j][i] == '#' {
				columnWithoutGalaxy = false
				break
			}
		}

		if columnWithoutGalaxy {
			emptyColumns = append(emptyColumns, i)
		}

	}

	for i := 0; i < len(universe); i++ {
		var rowWithoutGallaxy = true
		for j := 0; j < len(universe[i]); j++ {
			if universe[i][j] == '#' {
				rowWithoutGallaxy = false
				break
			}
		}
		newRow := make([]rune, len(emptyColumns)+len(universe[i]))

		var indexSwitch = 1
		for _, index := range emptyColumns {
			newRow[index+indexSwitch] = '.'
			indexSwitch++
		}

		for _, value := range universe[i] {
			var indexToPut = 0

			for {
				if newRow[indexToPut] == 0 {
					break
				}
				indexToPut++
			}

			newRow[indexToPut] = value
		}

		res = append(res, newRow)

		if rowWithoutGallaxy {

			res = append(res, newRow)

		}
	}

	return res
}

func getExpansionIndexes(universe [][]rune) (rows []int, cols []int) {

	for i := 0; i < len(universe[0]); i++ {
		var columnWithoutGalaxy = true
		for j := 0; j < len(universe); j++ {
			if universe[j][i] == '#' {
				columnWithoutGalaxy = false
				break
			}
		}

		if columnWithoutGalaxy {
			cols = append(cols, i)
		}
	}

	for i := 0; i < len(universe); i++ {
		var rowWithoutGallaxy = true
		for j := 0; j < len(universe[i]); j++ {
			if universe[i][j] == '#' {
				rowWithoutGallaxy = false
				break
			}
		}

		if rowWithoutGallaxy {
			rows = append(rows, i)
		}

	}

	return rows, cols
}

func ManhatanDistnace(a, b Coord, rowExpansions, colExpansions []int) float64 {
	var rowFactor int
	var colFactor int

	var xMin, xMax, yMin, yMax int

	if a.X > b.X {
		xMax = a.X
		xMin = b.X
	} else {
		xMax = b.X
		xMin = a.X
	}

	if a.Y > b.Y {
		yMax = a.Y
		yMin = b.Y
	} else {
		yMax = b.Y
		yMin = a.Y
	}

	for _, rowExpansion := range rowExpansions {
		if rowExpansion > yMin && rowExpansion < yMax {
			rowFactor++
		}
	}

	for _, colExpansion := range colExpansions {
		if colExpansion > xMin && colExpansion < xMax {
			colFactor++
		}
	}

	return math.Abs(float64(a.X)-float64(b.X)) + math.Abs(float64(a.Y)-float64(b.Y)) + (float64(rowFactor) * 999999) + (float64(colFactor) * 999999)
}

func main() {
	var universe [][]rune
	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	fileLines := strings.Split(string(file), "\n")

	for _, lines := range fileLines {
		universe = append(universe, []rune(lines))
	}

	var galaxiesIndex []Coord

	rows, cols := getExpansionIndexes(universe)

	fmt.Println(rows)
	fmt.Println(cols)

	for i, line := range universe {
		for j, value := range line {
			if value == '#' {
				galaxiesIndex = append(galaxiesIndex, Coord{X: j, Y: i})
			}
		}
	}

	var total int

	for i := 0; i < len(galaxiesIndex); i++ {
		for j := i; j < len(galaxiesIndex); j++ {
			if i != j {
				total += int(ManhatanDistnace(galaxiesIndex[i], galaxiesIndex[j], rows, cols))
			}
		}
	}

	fmt.Printf("total is: %d\n", total)

}
