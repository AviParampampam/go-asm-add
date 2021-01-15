package main

import "fmt"

// Run: go build -o add && ./add

func add1(x, y int64) int64

func main() {
	fmt.Println(add1(2, 3))
}
