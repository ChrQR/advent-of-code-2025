package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/gosuri/uilive"
)

type Range struct {
	start int
	end   int
}

func (r Range) Overlaps(c Range) bool {
	return r.start <= c.end && c.start <= r.end
}

func main() {

	writer := uilive.New()
	writer.Start()

	var ranges []Range

	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			continue
		}
		if strings.Contains(scanner.Text(), "-") {
			splitInput := strings.Split(scanner.Text(), "-")
			start, _ := strconv.Atoi(splitInput[0])
			end, _ := strconv.Atoi(splitInput[1])
			ranges = append(ranges, Range{
				start,
				end,
			})
		} else {
			break
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	merged := []Range{ranges[0]}
	for i := 1; i < len(ranges); i++ {
		last := &merged[len(merged)-1]
		if ranges[i].start <= last.end+1 {
			last.end = max(last.end, ranges[i].end)
		} else {
			merged = append(merged, ranges[i])
		}
	}

	freshIngredients := 0
	for _, r := range merged {
		freshIngredients += r.end - r.start + 1
		fmt.Fprintf(writer, "Added %d \n", r.end-r.start+1)
	}
	writer.Stop()
	fmt.Println(freshIngredients)
}
