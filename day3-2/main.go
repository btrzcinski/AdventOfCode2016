package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_possible_triangle(a, b, c int) bool {
	return a+b > c &&
		b+c > a &&
		a+c > b
}

func main() {
	input_file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer input_file.Close()

	possible_triangles := 0

	scanner := bufio.NewScanner(input_file)
	var dataset [3][]string
	for scanner.Scan() {
		dataset[0] = strings.Fields(scanner.Text())
		scanner.Scan()
		dataset[1] = strings.Fields(scanner.Text())
		scanner.Scan()
		dataset[2] = strings.Fields(scanner.Text())

		for i := 0; i < 3; i++ {
			a, _ := strconv.Atoi(dataset[0][i])
			b, _ := strconv.Atoi(dataset[1][i])
			c, _ := strconv.Atoi(dataset[2][i])

			if is_possible_triangle(a, b, c) {
				possible_triangles++
			}
		}
	}

	fmt.Printf("Possible triangles: %d\n", possible_triangles)
}
