package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func hasRepeatingPattern(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen == 0 {
			pattern := s[:patternLen]
			if strings.Repeat(pattern, length/patternLen) == s {
				return true
			}
		}
	}
	return false
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
			if hasRepeatingPattern(i) {
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
