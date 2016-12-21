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

type Int int

func (x Int) abs() Int {
	if x < 0 {
		return -x
	}
	return x
}

type position struct {
	x Int
	y Int
}

func (p *position) dist() Int {
	return p.x.abs() + p.y.abs()
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
	current_position := position{
		x: 0,
		y: 0,
	}
	dirs := []byte{'N', 'E', 'S', 'W'}
	facing := 0
	visited := make(map[position]bool)
MoveLoop:
	for _, move := range read_moves() {
		//fmt.Printf("%c: %d\n", move.direction, move.amount)
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

		for i := 0; i < move.amount; i++ {
			switch dirs[facing] {
			case 'N':
				current_position.y++
			case 'E':
				current_position.x++
			case 'S':
				current_position.y--
			case 'W':
				current_position.x--
			}

			if visited[current_position] {
				fmt.Printf("Visited twice: X: %d, Y: %d, distance: %d\n",
					current_position.x, current_position.y, current_position.dist())
				break MoveLoop
			}

			visited[current_position] = true
		}
	}
}
