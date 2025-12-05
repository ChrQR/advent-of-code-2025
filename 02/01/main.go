package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func hasRepeatingPatternTwice(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	if length%2 != 0 {
		return false
	}

	half := length / 2
	return s[:half] == s[half:]
}

func main() {
	data, err := os.ReadFile("./input.txt")
	fatalErr(err)

	counter := 0

	for input := range strings.SplitSeq(string(data), ",") {
		input = strings.TrimSpace(input)
		rangeList := strings.Split(input, "-")
		start, _ := strconv.Atoi(rangeList[0])
		end, _ := strconv.Atoi(rangeList[1])

		for i := start; i <= end; i++ {
			if hasRepeatingPatternTwice(i) {
				counter += i
			}
		}
	}
	fmt.Print(counter)
}
func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
