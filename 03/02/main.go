package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() {
		total += highestNumber(scanner.Text())
	}
	fmt.Print(total)
}

func highestNumber(input string) int {
	rslt := ""
	highest := 0
	highestIdx := 0
	i := 0
	for {
		num, _ := strconv.Atoi(string(input[i]))

		if num > highest {
			highest = num
			highestIdx = i
		}

		if i+(12-len(rslt)) == len(input) {
			rslt += string(input[highestIdx])
			i = highestIdx
			highest = 0
		}

		if len(rslt) == 12 {
			break
		}

		i++

	}
	rsltInt, _ := strconv.Atoi(rslt)
	fmt.Printf("Result: %d\n", rsltInt)
	fmt.Printf("Result len: %d \n", len(rslt))
	return rsltInt
}
