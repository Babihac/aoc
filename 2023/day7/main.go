package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
}

func CompareHands(handA, handB Hand, cardPowers map[rune]int) bool {
	mapA := make(map[rune]int)
	mapB := make(map[rune]int)

	var handAPower int
	var handBPower int

	handBChars := []rune(handB.cards)

	// It should no be same
	if handA == handB {
		return true
	}

	for _, char := range handA.cards {
		mapA[char]++
	}

	for _, char := range handB.cards {
		mapB[char]++
	}

	// if there are only J's (len == 1) treat J as J
	if mapA['J'] > 0 && len(mapA) > 1 {
		var maxValue = -1
		var maxKey rune
		for key, value := range mapA {
			if key == 'J' {
				continue
			}
			if value > maxValue {
				maxKey = key
				maxValue = value
			}
		}
		mapA[maxKey] += mapA['J']
		delete(mapA, 'J')

	}

	if mapB['J'] > 0 && len(mapB) > 1 {
		var maxValue = -1
		var maxKey rune
		for key, value := range mapB {
			if key == 'J' {
				continue
			}
			if value > maxValue {
				maxKey = key
				maxValue = value
			}
		}
		mapB[maxKey] += mapB['J']
		delete(mapB, 'J')

	}

	for _, value := range mapA {

		handAPower |= 1 << value
	}

	for _, value := range mapB {
		handBPower |= 1 << value
	}

	// A has better combination than B
	if handAPower > handBPower {
		return false
	}

	// B has better combination than A
	if handAPower < handBPower {
		return true
	}

	// both has same best combination but there might be two pairs or just one, just check the len of map
	if len(mapA) < len(mapB) {
		return false
	}

	if len(mapB) < len(mapA) {
		return true
	}

	// check which hand has stronger card in case of draw
	for i, char := range handA.cards {
		if cardPowers[char] == cardPowers[handBChars[i]] {
			continue
		}
		if cardPowers[char] > cardPowers[handBChars[i]] {
			return false
		} else {
			break
		}
	}
	return true
}

func main() {
	cardPowers := make(map[rune]int)
	cards := "J23456789TQKA"
	var hands []Hand

	for i, char := range cards {
		cardPowers[char] = i
	}

	file, err := os.Open("./input.txt")

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		handAndBid := strings.Fields(scanner.Text())

		if len(handAndBid) < 2 {
			panic("wrong input")
		}

		bid, err := strconv.Atoi(handAndBid[1])

		if err != nil {
			log.Fatal(err)
		}

		hands = append(hands, Hand{cards: handAndBid[0], bid: bid})

	}

	sort.Slice(hands, func(i, j int) bool {
		return CompareHands(hands[i], hands[j], cardPowers)
	})

	var mult = 1
	var result int

	for _, hand := range hands {
		result += hand.bid * mult
		mult++
	}

	fmt.Printf("The result is: %d\n", result)

}
