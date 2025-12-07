package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/gosuri/uilive"
)

func main() {
	var data [][]string

	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), ""))
	}

	writer := uilive.New()
	writer.Start()
	result := 0
	checkData(data, &result, writer)
	print(result)
}

func checkData(d [][]string, r *int, w *uilive.Writer) {
	output := ""
	movable := 0
	for y, v := range d {
		for x, k := range v {
			if k == "@" {
				occupied := 0
				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						if dy == 0 && dx == 0 {
							continue
						}
						ny, nx := y+dy, x+dx
						if ny >= 0 && ny < len(d) && nx >= 0 && nx < len(d[ny]) {
							if d[ny][nx] == "@" {
								occupied++
							}
						}
					}
				}
				if occupied < 4 {
					output += "."
					d[y][x] = "."
					movable++
					*r++
				} else {
					output += k
				}
			} else {
				output += k
			}
		}
		output += "\n"
	}

	fmt.Fprint(w, output)
	w.Flush()

	if movable > 0 {
		checkData(d, r, w)
	}
}
