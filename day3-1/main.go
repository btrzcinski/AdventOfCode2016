package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    input_file, err := os.Open("input.txt")
    if err != nil {
        panic(err)
    }
    defer input_file.Close()

    possible_triangles := 0

    scanner := bufio.NewScanner(input_file)
    scanner.Split(bufio.ScanWords)
    for scanner.Scan() {
        a, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        b, _ := strconv.Atoi(scanner.Text())
        scanner.Scan()
        c, _ := strconv.Atoi(scanner.Text())

        if a + b > c &&
           b + c > a &&
           a + c > b {
            possible_triangles += 1
        }
    }

    fmt.Printf("Possible triangles: %d\n", possible_triangles)
}
