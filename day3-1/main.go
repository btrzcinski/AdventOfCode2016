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
	for scanner.Scan() {
		triangle := strings.Fields(scanner.Text())

		a, _ := strconv.Atoi(triangle[0])
		b, _ := strconv.Atoi(triangle[1])
		c, _ := strconv.Atoi(triangle[2])

		if is_possible_triangle(a, b, c) {
			possible_triangles++
		}
	}

	fmt.Printf("Possible triangles: %d\n", possible_triangles)
}
