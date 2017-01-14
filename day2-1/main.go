package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	input_data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(strings.TrimSpace(string(input_data)), "\n")
	for _, line := range lines {
		x, y := 1, 1
		moves := strings.Split(line, "")
		for _, move := range moves {
			switch move {
			case "L":
				if x > 0 {
					x -= 1
				}
			case "U":
				if y > 0 {
					y -= 1
				}
			case "R":
				if x < 2 {
					x += 1
				}
			case "D":
				if y < 2 {
					y += 1
				}
			}
		}
		code := 1 + (3 * y) + x
		fmt.Printf("%d", code)
	}
	fmt.Printf("\n")
}
