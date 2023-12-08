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

type GameData struct {
	Id                 int
	ShownCubes         []Cubes
	fewestCubesForGame Cubes
}

type Cubes struct {
	Red   int
	Green int
	Blue  int
}

func setFewesCubesForGame(data *GameData, color string, count int) {
	switch color {
	case "red":
		data.fewestCubesForGame.Red = int(math.Max(float64(count), float64(data.fewestCubesForGame.Red)))
	case "green":
		data.fewestCubesForGame.Green = int(math.Max(float64(count), float64(data.fewestCubesForGame.Green)))
	case "blue":
		data.fewestCubesForGame.Blue = int(math.Max(float64(count), float64(data.fewestCubesForGame.Blue)))
	}
}

func setColor(c *Cubes, color string, count int) {
	switch color {
	case "red":
		c.Red = count
	case "green":
		c.Green = count
	case "blue":
		c.Blue = count
	}
}

func NewGameData(input string) *GameData {
	var gameData = GameData{}

	idAndGames := strings.Split(input, ":")
	cubesPerGame := strings.Split(idAndGames[1], ";")
	var gameCubes []Cubes
	countReg := regexp.MustCompile(`\d+`)
	colorReg := regexp.MustCompile(`[^\d\s]+`)
	for _, cubesInput := range cubesPerGame {
		var cubes Cubes
		currentGameCubes := strings.Split(cubesInput, ",")

		for _, cubesOfColor := range currentGameCubes {
			matches := countReg.FindString(cubesOfColor)
			color := colorReg.FindString(cubesOfColor)
			count, err := strconv.Atoi(matches)

			if err != nil {
				log.Fatal(err)
			}

			setColor(&cubes, color, count)
			setFewesCubesForGame(&gameData, color, count)
		}
		gameCubes = append(gameCubes, cubes)
	}
	gameId, err := strconv.Atoi(strings.TrimPrefix(idAndGames[0], "Game "))

	if err != nil {
		log.Fatal(err)
	}

	gameData.Id = gameId
	gameData.ShownCubes = gameCubes

	return &gameData
}

func main() {
	const NUM_OF_RED = 12
	const NUM_OF_GREEN = 13
	const NUM_OF_BLUE = 14

	var totalSumOfIds int
	var powerOfCubes int

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		gameData := NewGameData(scanner.Text())
		powerOfCubes += (gameData.fewestCubesForGame.Green * gameData.fewestCubesForGame.Red * gameData.fewestCubesForGame.Blue)
		var validGame bool = true

		for _, selectedCubes := range gameData.ShownCubes {
			redCubesLeft := NUM_OF_RED - selectedCubes.Red
			GreenCubesLeft := NUM_OF_GREEN - selectedCubes.Green
			blueCubesLeft := NUM_OF_BLUE - selectedCubes.Blue

			if redCubesLeft < 0 || GreenCubesLeft < 0 || blueCubesLeft < 0 {
				validGame = false
				break
			}
		}

		if validGame {
			totalSumOfIds += gameData.Id
		}

		fmt.Println(gameData)
	}

	fmt.Printf("Total sum of ids is: %d and power of cubes is: %d", totalSumOfIds, powerOfCubes)

}
