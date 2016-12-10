package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type move struct {
	direction byte
	amount    int
}

func read_moves() []move {
	input_data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	input_data_str := strings.Split(strings.TrimSpace(string(input_data)), ", ")

	moves := make([]move, len(input_data_str))
	for i := 0; i < len(input_data_str); i++ {
		move_amount, _ := strconv.Atoi(input_data_str[i][1:])
		moves[i] = move{
			direction: input_data_str[i][0],
			amount:    move_amount,
		}
	}

	return moves
}

func main() {
	x, y := 0, 0
	dirs := []byte{'N', 'E', 'S', 'W'}
	facing := 0
	for _, move := range read_moves() {
		fmt.Printf("%c: %d\n", move.direction, move.amount)
		switch move.direction {
		case 'R':
			facing++
		case 'L':
			facing--
		default:
			panic("Bad direction given")
		}

		if facing > 3 {
			facing = 0
		} else if facing < 0 {
			facing = 3
		}

		switch dirs[facing] {
		case 'N':
			y += move.amount
		case 'E':
			x += move.amount
		case 'S':
			y -= move.amount
		case 'W':
			x -= move.amount
		}
	}

	fmt.Printf("X: %d, Y: %d\n", x, y)
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	fmt.Printf("Distance: %d\n", x+y)
}
