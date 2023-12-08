package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Direction struct {
	left  string
	right string
}

func (d *Direction) ChooseDirection(dir rune) string {
	if dir == 'R' {
		return d.right
	} else {
		return d.left
	}
}

func walk(dirs []rune, moves map[string]Direction, start string) int {
	var totalSteps int
	var currentPosition = start
	var dirIndex int
	var dirsLen = len(dirs)

	for {

		// for part one this should be replaced by checking currentPosition == "ZZZ"
		if strings.HasSuffix(currentPosition, "Z") {
			break
		}

		currentDirection := moves[currentPosition]
		direction := dirs[dirIndex%dirsLen]

		currentPosition = currentDirection.ChooseDirection(direction)

		dirIndex++
		totalSteps++

	}
	return totalSteps
}

func main() {
	var dirs []rune
	starts := make([]string, 0, 10)
	moves := make(map[string]Direction)

	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(file), "\n")

	dirs = []rune(data[0])

	for _, line := range data[2:] {
		entries := strings.Split(line, "=")

		key := strings.TrimSpace(entries[0])

		directions := strings.Split(strings.Trim(entries[1], " ()"), ", ")

		moves[key] = Direction{left: directions[0], right: directions[1]}

		if strings.HasSuffix(key, "A") {
			starts = append(starts, key)
		}

	}

	// calculate LCM of end numbers - I am lazy to implement it here so i just used python code from CHGPT...
	// 	from math import gcd
	// from functools import reduce

	// def lcm(a, b):
	//     return a * b // gcd(a, b)

	// def find_lcm(numbers):
	//     return reduce(lcm, numbers)

	// # Given numbers
	// numbers = [16343, 11911, 20221, 21883, 13019, 19667] // these are the steps for each end

	// # Calculate LCM
	// result = find_lcm(numbers)

	// # Output the result
	// print(f"The LCM of the given numbers is: {result}")

	for _, start := range starts {
		res := walk(dirs, moves, start)
		fmt.Println(res)
	}
}
