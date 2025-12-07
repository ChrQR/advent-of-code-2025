package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func (r Range) Contains(n int) bool {
	if r.start <= n && n <= r.end {
		return true
	}
	return false
}

func main() {
	var ranges []Range
	var ingredients []int

	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "-") {
			ranges = append(ranges, toRange(scanner.Text()))
		} else if scanner.Text() != "" {
			ingr, _ := strconv.Atoi(scanner.Text())
			ingredients = append(ingredients, ingr)
		}
	}

	fresh := 0

	for _, ingredient := range ingredients {
		for _, r := range ranges {
			if r.Contains(ingredient) {
				fresh++
				break
			}
		}
	}
	fmt.Print(fresh)
}

func toRange(r string) Range {
	splitStr := strings.Split(r, "-")
	start, _ := strconv.Atoi(splitStr[0])
	end, _ := strconv.Atoi(splitStr[1])
	return Range{
		start,
		end,
	}
}
