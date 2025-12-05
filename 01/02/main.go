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
	fatalErr(err)

	zeroCounter := 0

	for input := range strings.SplitSeq(string(inputFile), "\n") {
		if len(input) == 0 {
			continue
		}
		var direction byte
		var num int
		direction = input[0]
		parsedNum, err := strconv.Atoi(input[1:])
		fatalErr(err)

		if len(input) > 3 {
			firstDigits, err := strconv.Atoi(input[1 : len(input)-2])
			zeroCounter += firstDigits - 1
			fatalErr(err)
		}

		num = parsedNum % 100

		switch string(direction) {
		case "L":
			dial -= num
		case "R":
			dial += num
		default:
			log.Fatal("No direction detected!")
		}

		if dial < 0 {
			zeroCounter++
			dial += 100
		}
		if dial > 99 {
			zeroCounter++
			dial -= 100
		}

		if dial == 0 {
			zeroCounter++
		}
	}

	fmt.Printf("Total 0s: %d", zeroCounter)
}

func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
