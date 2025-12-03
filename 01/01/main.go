package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	dial := 50

	inputFile, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	zeroCounter := 0

	for input := range strings.SplitSeq(string(inputFile), "\n") {
		if len(input) == 0 {
			continue
		}
		var direction byte
		var num int
		direction = input[0]
		parsedNum, err := strconv.Atoi(input[1:])
		num = parsedNum % 100
		if err != nil {
			log.Fatalf("unable to parse string to int: %v", err)
		}

		switch string(direction) {
		case "L":
			dial -= num
		case "R":
			dial += num
		default:
			log.Fatal("No direction detected!")
		}

		if dial < 0 {
			dial += 100
		}
		if dial > 99 {
			dial -= 100
		}

		if dial == 0 {
			zeroCounter++
		}
	}

	fmt.Printf("Total 0s: %d", zeroCounter)
}
