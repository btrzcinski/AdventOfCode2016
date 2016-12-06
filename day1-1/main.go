package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func read_moves() []string {
	input_data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input_data_str := strings.TrimSpace(string(input_data))

	return strings.Split(input_data_str, ", ")
}

func main() {
	for _, move := range read_moves() {
		fmt.Printf("%s\n", move)
	}
}
