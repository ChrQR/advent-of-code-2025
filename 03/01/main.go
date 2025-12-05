package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	total := MaxJolt(scanner)

	fmt.Print(total)
}

func MaxJolt(scanner *bufio.Scanner) int {

	total := 0

	for scanner.Scan() {
		first := 0
		firstIdx := 0
		second := 0
		for i, char := range scanner.Text() {
			num, _ := strconv.Atoi(string(char))
			if num > first && i != len(scanner.Text())-1 {
				first = num
				firstIdx = i
			}
		}
		for _, scndChar := range scanner.Text()[firstIdx+1:] {
			scndNum, _ := strconv.Atoi(string(scndChar))
			if scndNum > second {
				second = scndNum
			}

		}
		totalStr := strconv.Itoa(first)
		totalStr += strconv.Itoa(second)
		sum, _ := strconv.Atoi(totalStr)
		log.Print(sum)
		total += sum
	}
	return total
}
