package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func nextSequence(sequence []int) ([]int, bool) {
	var res []int
	var allZeroDiff = true
	for i := 0; i < len(sequence)-1; i++ {
		var diff = sequence[i+1] - sequence[i]

		if diff != 0 {
			allZeroDiff = false
		}

		res = append(res, diff)
	}
	return res, allZeroDiff
}

func readLine(line string) []int {
	var res []int

	values := strings.Fields(line)

	for _, value := range values {
		num, err := strconv.Atoi(value)

		if err != nil {
			log.Fatal(err)
		}

		res = append(res, num)
	}
	return res
}

func extrapolate(seq []int) int {
	var stack []int

	for {
		stack = append(stack, seq[0])

		nextSeq, allZero := nextSequence(seq)

		if allZero {
			break
		}
		seq = nextSeq
	}

	var extrapolatedValue int

	for {
		if len(stack) == 0 {
			break
		}

		value := stack[len(stack)-1]
		extrapolatedValue = value - extrapolatedValue

		stack = stack[:len(stack)-1]

	}
	return extrapolatedValue
}

func main() {
	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var res int

	for scanner.Scan() {
		sequence := readLine(scanner.Text())

		res += extrapolate(sequence)
	}

	fmt.Printf("The result is: %d\n", res)

}
