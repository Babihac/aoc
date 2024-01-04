package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func hashResult(input string) int {

	var currentValue int

	for _, char := range input {
		asciiValue := int(char)
		currentValue = ((currentValue + asciiValue) * 17) % 256

	}

	return currentValue

}

func main() {

	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	data := strings.Split(string(file), ",")

	var result int

	start := time.Now()

	for _, input := range data {

		hashResult := hashResult(input)
		result += hashResult
	}

	elapsed := time.Since(start)

	fmt.Printf("The result is %d\n", result)
	fmt.Printf("The program took %s to run.\n", elapsed)

}
