package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/gosuri/uilive"
)

func main() {
	var data [][]string

	file, _ := os.Open("./input.txt")

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		data = append(data, strings.Split(scanner.Text(), ""))
	}
	movable := checkData(data)
	print(movable)
}

func checkData(d [][]string) int {
	writer := uilive.New()
	writer.Start()
	output := ""
	movable := 0

	for y, v := range d {
		for x, k := range v {
			output += k

			if k == "@" {
				occupied := 0

				for dy := -1; dy <= 1; dy++ {
					for dx := -1; dx <= 1; dx++ {
						// Skip the center cell
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
					output = output[:len(output)-1] + color.GreenString("")
					movable++
				}
			}
		}
		output += "\n"
		fmt.Fprint(writer, output)
	}

	return movable
}
