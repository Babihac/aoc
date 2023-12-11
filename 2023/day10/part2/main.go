package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readLine(line string) []rune {
	var res []rune

	for _, char := range line {
		if char == '.' {
			res = append(res, '0')
		} else {
			res = append(res, '1')
		}
	}
	return res
}

func main() {
	file, err := os.Open("./input.txt")

	var nodes [][]rune

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		nodes = append(nodes, readLine(scanner.Text()))

	}
	for _, line := range nodes {
		for _, node := range line {
			fmt.Printf("%c ", node)
		}
		fmt.Println()
	}
}
