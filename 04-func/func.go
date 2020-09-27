package main

import "fmt"

func add(x, y int) int {
	return x + y
}

// 等於 add(x int, y int) int {}

func main() {
	fmt.Println(add(1, 2))
}
