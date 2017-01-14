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
    grid := [][]int{
        {0, 0, 1, 0, 0},
        {0, 2, 3, 4, 0},
        {5, 6, 7, 8, 9},
        {0, 10, 11, 12, 0},
        {0, 0, 13, 0, 0},
    }
    for _, line := range lines {
        x, y := 0, 3
        moves := strings.Split(line, "")
        for _, move := range moves {
            switch move {
            case "L":
                if x > 0 && grid[y][x-1] > 0 {
                    x -= 1
                }
            case "U":
                if y > 0 && grid[y-1][x] > 0 {
                    y -= 1
                }
            case "R":
                if x < 4 && grid[y][x+1] > 0 {
                    x += 1
                }
            case "D":
                if y < 4 && grid[y+1][x] > 0 {
                    y += 1
                }
            }
        }
        code := grid[y][x]
        fmt.Printf("%X", code)
    }
    fmt.Printf("\n")
}
